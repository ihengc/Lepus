package acceptor

import (
	"net"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 15:20
* @version: 1.0
* @description:
*********************************************************/

type UDPAcceptor struct {
}

func (upd *UDPAcceptor) Run() {
	//TODO implement me
	panic("implement me")
}

func (upd *UDPAcceptor) Stop() {
	//TODO implement me
	panic("implement me")
}

func (upd *UDPAcceptor) LocalAddr() string {
	//TODO implement me
	panic("implement me")
}

func (upd *UDPAcceptor) GetConnChan() chan net.Conn {
	//TODO implement me
	panic("implement me")
}
