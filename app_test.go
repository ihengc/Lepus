package Lepus

import (
	"Lepus/acceptor"
	"sync"
	"testing"
	"time"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/21 0021 14:50
* @version: 1.0
* @description:
*********************************************************/

var (
	app         *Application
	tcpAcceptor *acceptor.TCPAcceptor
)

func newApp() {
	app = NewApplication("Test")
}

func newTcpAcceptor() {
	tcpAcceptor = acceptor.NewTCPAcceptor("localhost", 9017)
}

func setUp() {
	newApp()
	newTcpAcceptor()
	app.RegisterAcceptor(tcpAcceptor)
}

func TestApplication_Run(t *testing.T) {
	setUp()
	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		app.Run()
		group.Done()
	}()
	go func() {
		time.Sleep(time.Second * 3)
		app.Stop()
		app.Stop()
		group.Done()
	}()
	group.Wait()
}
