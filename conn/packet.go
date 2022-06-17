package conn

import (
	"encoding/binary"
	"io"
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

func DecodePacket(conn io.Reader) (*Packet, error) {
	buffer := make([]byte, 13)
	_, err := io.ReadFull(conn, buffer)
	if err != nil {
		return nil, err
	}

	serverId := binary.LittleEndian.Uint32(buffer[:4])
	serverType := binary.LittleEndian.Uint32(buffer[4:8])
	packetType := buffer[8:9]
	packetSize := binary.LittleEndian.Uint32(buffer[9:13])

	packet := &Packet{}

	packet.ServerId = serverId
	packet.ServerType = serverType
	packet.PacketType = PacketType(packetType[0])
	packet.PacketSize = packetSize

	return packet, nil
}

func EncodePacket(packet *Packet) []byte {
	buffer := make([]byte, packet.PacketSize+13)

	binary.BigEndian.PutUint32(buffer, packet.ServerId)

	binary.BigEndian.PutUint32(buffer, packet.ServerType)

	buffer = append(buffer, uint8(packet.PacketType))

	binary.BigEndian.PutUint32(buffer, packet.PacketSize)

	buffer = append(buffer, packet.Data...)

	return buffer
}
