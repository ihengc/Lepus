package Lepus

import (
	"Lepus/acceptor"
	"Lepus/codec"
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

type IApplication interface {
	Run()                                // 启动应用
	Stop()                               // 停止应用
	GetAppName() string                  // 获取应用名称
	RegisterAcceptor(acceptor.IAcceptor) // 配置监听对象
}

type AppOpt func(*Application)

func SetPacketCodec(packetCodec codec.IPacketCodec) AppOpt {
	return func(app *Application) {
		app.packetCodec = packetCodec
	}
}

// Application 实现 IApplication
type Application struct {
	name           string               // 应用名称
	isRunning      bool                 // 标识应用是否正在运行
	closeChan      chan bool            // 服务停止通道
	acceptors      []acceptor.IAcceptor // 监听对象
	packetCodec    codec.IPacketCodec   // 数据包编解码器
	handlerService *service.HandlerService
}

// runAcceptor 运行acceptor
func (app *Application) runAcceptor() {
	logger.Log("runAcceptor")
	for _, apt := range app.acceptors {
		go apt.Run()

		go func(apt acceptor.IAcceptor) {
			connChan := apt.GetConnChan()
			app.handlerService.Handle(connChan)
		}(apt)
	}
}

// Run 运行应用
func (app *Application) Run() {
	app.isRunning = true
	app.runAcceptor()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-signalChan:
		app.Stop()
	case <-app.closeChan:
		return
	}
}

// GetAppName 获取应用名称
func (app *Application) GetAppName() string {
	return app.name
}

// RegisterAcceptor 注册监听
func (app *Application) RegisterAcceptor(iAcceptor acceptor.IAcceptor) {
	app.acceptors = append(app.acceptors, iAcceptor)
}

// Stop 停止应用
func (app *Application) Stop() {
	select {
	case <-app.closeChan:
	default:
		app.isRunning = false
		app.closeChan <- true
		close(app.closeChan)
	}
}

// NewApplication 创建应用
func NewApplication(name string) *Application {
	app := new(Application)
	app.name = name
	app.closeChan = make(chan bool)
	return app
}

// NewApplicationWithOption 创建应用
func NewApplicationWithOption(opts ...AppOpt) *Application {
	app := new(Application)
	for _, opt := range opts {
		opt(app)
	}
	return app
}
