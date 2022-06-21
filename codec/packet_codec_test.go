package codec

import (
	"Lepus/conn"
	"bytes"
	"testing"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/21 0021 17:24
* @version: 1.0
* @description:
*********************************************************/

var (
	codec  *DefaultPacketCodec
	packet *conn.DefaultPacket
	data   []byte
)

func newCodec() {
	codec = new(DefaultPacketCodec)
}

func newPacket() {
	packet = new(conn.DefaultPacket)
	packet.SetID(1)
	packet.SetData([]byte("Test"))
}

func setup() {
	newCodec()
	newPacket()
}

func TestDefaultPacketCodec_Encode(t *testing.T) {
	setup()
	var err error
	data, err = codec.Encode(packet)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDefaultPacketCodec_Decode(t *testing.T) {
	reader := bytes.NewBuffer(data)
	packet, err := codec.Decode(reader)
	if err != nil {
		t.Fatal(err)
	}
	if packet.GetID() != 1 || string(packet.GetData()) != "Test" {
		t.Log("codec decode error")
	}
}
