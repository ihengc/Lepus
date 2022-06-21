package service

import (
	"Lepus/codec"
	connPkg "Lepus/conn"
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

// HandlerService 连接服务
type HandlerService struct {
	packetCodec codec.IPacketCodec
}

// Handle 处理连接
func (h *HandlerService) Handle(connChan chan net.Conn) {
	for conn := range connChan {
		packet, err := h.packetCodec.Decode(conn)
		if err != nil {
			logger.Err(err)
			continue
		}
		h.parsePacket(packet)
	}
}

// parsePacket 分析包
func (h *HandlerService) parsePacket(packet connPkg.IPacket) {
	fmt.Println(packet.GetID())
}

// NewHandlerService 创建连接处理服务
func NewHandlerService(packetCodec codec.IPacketCodec) *HandlerService {
	h := new(HandlerService)
	h.packetCodec = packetCodec
	return h
}
