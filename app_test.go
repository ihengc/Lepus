package Lepus

import (
	"Lepus/acceptor"
	"net"
	"sync"
	"testing"
	"time"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 16:18
* @version: 1.0
* @description:
*********************************************************/

var app *Application

func init() {
	app = NewApplication()
	tcp := acceptor.NewTCPAcceptor("localhost", 9017)
	app.RegisterAcceptor(tcp)
}

func TestApplication_Run(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		app.Run()
		group.Done()
	}()
	go func(t *testing.T) {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			conn, err := net.Dial("tcp", "localhost:9017")
			if err != nil {
				t.Error(err)
				return
			}
			t.Log("send message count:", i+1)
			_, err = conn.Write([]byte("ping"))
			if err != nil {
				t.Error(err)
				return
			}
		}
		app.Stop()
		group.Done()
	}(t)
	group.Wait()
}
