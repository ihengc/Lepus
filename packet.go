package Lepus

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 16:30
* @version: 1.0
* @description:
*********************************************************/

type PacketType byte

type Packet struct {
	Type PacketType
	Data []byte
}

type TCPPacket struct {
	Type PacketType
	Data []byte
}
