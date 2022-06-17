package Lepus

import (
	"Lepus/acceptor"
	"net"
	"sync"
	"testing"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/17 21:25
 * @description:
 ***************************************************************/

var app *Application

func init() {
	app = NewApplication()
	tcpAcceptor := acceptor.NewTCPAcceptor("localhost", 9017)
	app.RegisterAcceptor(tcpAcceptor)
}

func TestApplication_Run(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		app.Run()
		group.Done()
	}()
	go func(t *testing.T) {
		conn, err := net.Dial("tcp", "localhost:9017")
		if err != nil {
			t.Error(err.Error())
		}
		_, err = conn.Read([]byte("123"))
		if err != nil {
			return
		}
	}(t)
	group.Wait()
}
