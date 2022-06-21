package Lepus

import (
	"Lepus/acceptor"
	"Lepus/codec"
	"Lepus/conn"
	"encoding/binary"
	"io"
	"net"
	"sync"
	"testing"
	"time"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/21 0021 14:50
* @version: 1.0
* @description:
*********************************************************/

var (
	app         *Application
	tcpAcceptor *acceptor.TCPAcceptor
	pCodec      codec.IPacketCodec
)

type testPacket struct {
	ID   uint32
	Data []byte
}

func (t *testPacket) GetID() uint32 {
	return t.ID
}

func (t *testPacket) GetData() []byte {
	return t.Data
}

func (t *testPacket) SetID(id uint32) {
	t.ID = id
}

func (t *testPacket) SetData(data []byte) {
	t.Data = data
}

func (t *testPacket) GetLen() uint32 {
	return 8
}

type testPacketCodec struct{}

func (p *testPacketCodec) Encode(packet conn.IPacket) ([]byte, error) {
	idBuffer := make([]byte, 4) // packet id
	binary.BigEndian.PutUint32(idBuffer, packet.GetID())

	lenBuffer := make([]byte, 4) // packet size
	binary.BigEndian.PutUint32(lenBuffer, uint32(len(packet.GetData())))

	idBuffer = append(idBuffer, lenBuffer...)
	idBuffer = append(idBuffer, packet.GetData()...)
	return idBuffer, nil
}

func (p *testPacketCodec) Decode(reader io.Reader) (conn.IPacket, error) {
	pk := &testPacket{}

	idBuffer := make([]byte, 4)
	if _, err := io.ReadFull(reader, idBuffer); err != nil {
		id := binary.BigEndian.Uint32(idBuffer)
		pk.SetID(id)
	}

	lenBuffer := make([]byte, 4)
	if _, err := io.ReadFull(reader, lenBuffer); err != nil {
		size := binary.BigEndian.Uint32(lenBuffer)
		data := make([]byte, size)
		if _, err := io.ReadFull(reader, data); err != nil {
			pk.SetData(data)
		}
	}
	return pk, nil
}

func newPacketCodec() {
	pCodec = &testPacketCodec{}
}

func newApp() {
	newPacketCodec()
	app = NewApplicationWithOption("Test", SetPacketCodec(pCodec))
}

func newTcpAcceptor() {
	tcpAcceptor = acceptor.NewTCPAcceptor("localhost", 9017)
}

func setUp() {
	newApp()
	newTcpAcceptor()
	app.RegisterAcceptor(tcpAcceptor)
}

func sendMsg(t *testing.T) {
	c := new(testPacketCodec)
	p := new(testPacket)
	p.ID = 1
	p.Data = []byte("Test")
	data, err := c.Encode(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
	cn, err := net.Dial("tcp", "localhost:9017")
	if err != nil {
		t.Fatal(err)
	}
	_, err = cn.Write(data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApplication_Run(t *testing.T) {
	setUp()
	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		app.Run()
		group.Done()
	}()
	go func() {
		time.Sleep(time.Second * 3)
		//app.Stop()
		//app.Stop()
		sendMsg(t)
		group.Done()
	}()
	group.Wait()
}
