package acceptor

import (
	"sync"
	"testing"
	"time"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 15:31
* @version: 1.0
* @description:
*********************************************************/

var tcpAcceptor *TCPAcceptor

func init() {
	tcpAcceptor = NewTCPAcceptor("localhost", 9017)
}

func TestTCPAcceptor_Run(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(2)

	go func() {
		tcpAcceptor.Run()
		group.Done()
	}()
	go func() {
		time.Sleep(1 * time.Second)
		tcpAcceptor.Stop()
		time.Sleep(1 * time.Second)
		tcpAcceptor.Stop()
	}()
	group.Wait()
}
