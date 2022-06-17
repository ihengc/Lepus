package conn

import (
	"bytes"
	"testing"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/17 21:42
 * @description:
 ***************************************************************/

var packet *Packet

var buffer []byte

func init() {
	packet = &Packet{}
}

func TestEncodePacket(t *testing.T) {

	packet.PacketType = Heartbeat
	packet.ServerId = 1
	packet.ServerType = 10001
	packet.Data = []byte("Hello Lepus!!!")
	packet.PacketSize = uint32(len(packet.Data) + 13)

	buffer = EncodePacket(packet)

}

func TestDecodePacket(t *testing.T) {
	reader := bytes.NewReader(buffer)
	packet, err := DecodePacket(reader)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(packet.PacketType)
	t.Log(packet.ServerId)
}
