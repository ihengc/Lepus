package conn

import (
	"encoding/binary"
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

type Packet struct {
	ServerId   uint32
	ServerType uint32
	PacketType PacketType
	PacketSize uint32
	Data       []byte
}

func DecodePacket(conn net.Conn) (*Packet, error) {
	packetHeader := make([]byte, 9)
	_, err := io.ReadFull(conn, packetHeader)
	if err != nil {
		return nil, err
	}
	serverId := binary.LittleEndian.Uint32(packetHeader[:4])
	serverType := binary.LittleEndian.Uint32(packetHeader[4:8])
	packetType := binary.LittleEndian.Uint16(packetHeader[8:9])
	packetSize := binary.LittleEndian.Uint32(packetHeader[9:13])

	packet := &Packet{}

	packet.ServerId = serverId
	packet.ServerType = serverType
	packet.PacketType = PacketType(packetType)
	packet.PacketSize = packetSize

	return packet, nil
}
