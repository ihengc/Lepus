package acceptor

import (
	"Lepus/logger"
	"fmt"
	"net"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/21 0021 15:19
* @version: 1.0
* @description:
*********************************************************/

// TCPAcceptor tcp监听
type TCPAcceptor struct {
	ln        net.Listener
	isRunning bool
	connChan  chan net.Conn
}

// Run 启动监听
func (tcpAcceptor *TCPAcceptor) Run() {
	logger.Log("TCP Acceptor Run")
	defer tcpAcceptor.Stop()
	tcpAcceptor.isRunning = true

	for tcpAcceptor.isRunning {
		conn, err := tcpAcceptor.ln.Accept()
		if err != nil {
			logger.Err(err)
			continue
		}
		tcpAcceptor.connChan <- conn
	}
}

// Stop 停止监听
func (tcpAcceptor *TCPAcceptor) Stop() {
	err := tcpAcceptor.ln.Close()
	if err != nil {
		logger.Err(err)
	}
	tcpAcceptor.isRunning = false
}

// LocalAddr 获取监听地址
func (tcpAcceptor *TCPAcceptor) LocalAddr() string {
	return tcpAcceptor.ln.Addr().String()
}

// GetConnChan 获取连接队列
func (tcpAcceptor *TCPAcceptor) GetConnChan() chan net.Conn {
	return tcpAcceptor.connChan
}

// NewTCPAcceptor 创建TCP监听
func NewTCPAcceptor(host string, port int) *TCPAcceptor {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logger.Err(err)
		return nil
	}
	tcpAcceptor := new(TCPAcceptor)
	tcpAcceptor.ln = ln
	tcpAcceptor.connChan = make(chan net.Conn)
	return tcpAcceptor
}
