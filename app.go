package Lepus

import (
	"Lepus/acceptor"
	"Lepus/logger"
	"Lepus/service"
	"os"
	"os/signal"
	"syscall"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 10:37
* @version: 1.0
* @description:
*********************************************************/

type AppMode byte

const (
	Front AppMode = iota + 1
	Backend
)

type Application struct {
	name           string
	mode           AppMode
	running        bool
	closeChan      chan bool
	acceptor       *acceptor.TCPAcceptor
	handlerService *service.HandlerService
}

func (app *Application) runAcceptor() {
	if app.mode == Front {
		if app.acceptor == nil {
			logger.Log("application is front, but the acceptor is nil")
			return
		}
		go app.acceptor.Run()
		connChan := app.acceptor.GetConnChan()
		go app.handlerService.Handle(connChan)
	}

}

func (app *Application) Run() {
	app.running = true
	app.runAcceptor()
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	select {
	case <-signalChan:
		app.Stop()
	case <-app.closeChan:
		return
	}
}

func (app *Application) Stop() {
	select {
	case <-app.closeChan:
	default:
		app.running = false
		app.closeChan <- false
		close(app.closeChan)
	}
}

func (app *Application) RegisterAcceptor(tcpAcceptor *acceptor.TCPAcceptor) {
	app.acceptor = tcpAcceptor
}

func NewApplication() *Application {
	app := &Application{}
	app.name = "Lepus"
	app.closeChan = make(chan bool)
	app.running = false
	app.mode = Front
	return app
}
