package acceptor

import (
	"fmt"
	"log"
	"net"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/7/3 22:45
 * @description:
 ***************************************************************/

// TODO Golang中创建chan时make(chan x)和make(chan x, 1)在容量上有什么区别？

type TCPAcceptor struct {
	// address 监听的地址
	address string
	// ln 监听
	ln net.Listener
	// running 接收器是否运行
	running bool
	// connChan 用于传递连接
	connChan chan net.Conn
}

func (t *TCPAcceptor) listen() {
	ln, err := net.Listen("tcp", t.address)
	if err != nil {
		log.Fatalln("tcp acceptor listen err:", err)
		return
	}
	t.ln = ln
}

// Accept 接收连接
func (t *TCPAcceptor) Accept() {
	t.listen()
	t.running = true
	for t.running {
		conn, err := t.ln.Accept()
		if err != nil {
			continue
		}
		t.connChan <- conn
	}
}

// GetConnChan 获取连接chan
func (t *TCPAcceptor) GetConnChan() chan net.Conn {
	return t.connChan
}

// Address 获取监听地址
func (t *TCPAcceptor) Address() string {
	return t.address
}

// Stop 关闭接收器
func (t *TCPAcceptor) Stop() {
	t.running = false
	err := t.ln.Close()
	if err != nil {
		log.Println(err)
	}
}

// NewTCPAcceptor 创建TCP连接接收器
func NewTCPAcceptor(host string, port int) *TCPAcceptor {
	return &TCPAcceptor{
		address:  fmt.Sprintf("%s:%d", host, port),
		connChan: make(chan net.Conn),
	}
}
