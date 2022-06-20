package conn

import (
	"io"
	"net"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/17 20:26
 * @description:
 ***************************************************************/

type PacketType byte

const (
	Heartbeat PacketType = iota + 1
	Handshake
	Data
)

const PacketHeaderSize = 12

type Packet struct {
	ServerId   uint32
	ServerType uint32
	PacketType PacketType
	PacketSize uint32
	Data       []byte
}

func DecodePacket(conn net.Conn) (*Packet, error) {
	buffer := make([]byte, PacketHeaderSize)
	_, err := io.ReadFull(conn, buffer)
	if err != nil {
		return nil, err
	}
	packet := &Packet{}
	return packet, nil
}

func EncodePacket(packet *Packet) []byte {
	buffer := make([]byte, packet.PacketSize)
	return buffer
}
