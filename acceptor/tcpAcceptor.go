package acceptor

import (
	"Lepus/logger"
	"fmt"
	"net"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 11:50
* @version: 1.0
* @description:
*********************************************************/

type TCPAcceptor struct {
	running  bool
	connChan chan net.Conn // connChan 是无缓冲的
	listen   net.Listener
}

// Run 接入tcp连接
func (tcp *TCPAcceptor) Run() {
	defer tcp.Stop()
	tcp.running = true
	for tcp.running {
		conn, err := tcp.listen.Accept()
		if err != nil {
			logger.Log(fmt.Sprintf("tcp acceptor accept err:%s", err.Error()))
			continue
		}
		tcp.connChan <- conn
	}
}

// Stop 关闭监听
func (tcp *TCPAcceptor) Stop() {
	tcp.running = false
	err := tcp.listen.Close()
	if err != nil {
		logger.Log(fmt.Sprintf("tcp acceptor stop err:%s", err.Error()))
	}
}

func (tcp *TCPAcceptor) LocalAddr() string {
	return tcp.listen.Addr().String()
}

func (tcp *TCPAcceptor) GetConnChan() chan net.Conn {
	return tcp.connChan
}

func NewTCPAcceptor(host string, port int) *TCPAcceptor {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logger.Log(fmt.Sprintf("new tcp acceptor err:%s", err.Error()))
		return nil
	}
	tcp := &TCPAcceptor{}
	tcp.listen = listen
	tcp.running = false
	tcp.connChan = make(chan net.Conn, 0)
	return tcp
}
