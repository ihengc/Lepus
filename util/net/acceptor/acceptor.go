package acceptor

import "net"

/****************************************************************
 * @author: Ihc
 * @date: 2022/7/3 23:13
 * @description:
 ***************************************************************/

type IAcceptor interface {
	Address() string
	GetConnChan() chan net.Conn
}
