package service

import (
	connPackage "Lepus/conn"
	"Lepus/logger"
	"fmt"
	"net"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 15:54
* @version: 1.0
* @description:
*********************************************************/

type HandlerService struct {
}

func (h *HandlerService) Handle(connChan chan net.Conn) {
	for conn := range connChan {
		packet, err := connPackage.DecodePacket(conn)
		if err != nil {
			logger.Log(fmt.Sprintf("decode packet err:%s", err.Error()))
			continue
		}
		h.Dispatch(packet)
	}
}
func (h *HandlerService) Dispatch(packet *connPackage.Packet) {
	switch packet.PacketType {
	case connPackage.Handshake:
	case connPackage.Heartbeat:

	}
}
