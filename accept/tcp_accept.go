package accept

import (
	"net"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 11:50
* @version: 1.0
* @description:
*********************************************************/

type TCPAccept struct {
	ln       net.Listener
	running  bool
	connChan chan net.Conn
}

func (tcpAccept *TCPAccept) Run() {
	tcpAccept.running = true
	for tcpAccept.running {
		conn, err := tcpAccept.ln.Accept()
		if err != nil {
			continue
		}
		tcpAccept.connChan <- conn
	}
}

func (tcpAccept *TCPAccept) Stop() {
	tcpAccept.running = false
	tcpAccept.ln.Close()
	close(tcpAccept.connChan)
}

func (tcpAccept *TCPAccept) LocalAddr() string {
	return tcpAccept.ln.Addr().String()
}

func (tcpAccept *TCPAccept) GetConnChan() chan net.Conn {
	return tcpAccept.connChan
}
