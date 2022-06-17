package accept

import "net"

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 11:50
* @version: 1.0
* @description:
*********************************************************/

type IAccept interface {
	Run()
	Stop()
	LocalAddr() string
	GetConnChan() chan net.Conn
}
