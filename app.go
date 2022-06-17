package Lepus

import (
	"Lepus/acceptor"
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

type IApplication interface {
	Run()
	Stop()
}

type AppMode byte

const (
	Front AppMode = iota + 1
	Backend
)

type Application struct {
	name      string
	appMode   AppMode
	running   bool
	closeChan chan bool
	broker    *Broker
	accepts   []acceptor.IAcceptor // accepts 用于前台服务
}

func (app *Application) runAccepts() {
	for i := 0; i < len(app.accepts); i++ {
		apt := app.accepts[i]
		go apt.Run()
		connChan := apt.GetConnChan()
		go app.broker.Handle(connChan)
	}
}

func (app *Application) listenExitSignal() chan os.Signal {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	return signalChan
}

func (app *Application) Run() {
	app.running = true

	app.runAccepts()
	signalChan := app.listenExitSignal()

	select {
	case <-signalChan:
		app.Stop()
	case <-app.closeChan:
		return
	}
}

// Stop 停止应用
func (app *Application) Stop() {
	app.running = false
	select {
	case <-app.closeChan:
	default:
		app.closeChan <- false
		close(app.closeChan)
	}
}

func (app *Application) RegisterAcceptor(iAccept acceptor.IAcceptor) {
	app.accepts = append(app.accepts, iAccept)
}

func NewApplication(options ...Option) *Application {
	conf := &ApplicationConf{
		Name:    "Lepus",
		Host:    "localhost",
		Port:    9017,
		AppMode: Front,
	}
	for _, option := range options {
		option(conf)
	}
	app := &Application{}
	app.running = false
	app.appMode = conf.AppMode
	app.name = conf.Name
	app.closeChan = make(chan bool)
	app.accepts = make([]acceptor.IAcceptor, 0)
	return app
}
