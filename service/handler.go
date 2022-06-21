package service

import (
	"Lepus/codec"
	"net"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 15:54
* @version: 1.0
* @description:
*********************************************************/

type HandlerService struct {
	packetCodec codec.IPacketCodec
}

func (h *HandlerService) Handle(connChan chan net.Conn) {

}
