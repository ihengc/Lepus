package codec

import (
	"Lepus/conn"
	"encoding/binary"
	"io"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/21 0021 09:45
* @version: 1.0
* @description:
*********************************************************/

// PacketEncoder 数据包编码器
type PacketEncoder interface {
	Encode(conn.IPacket) ([]byte, error)
}

// PacketDecoder 数据包解码器
type PacketDecoder interface {
	Decode(io.Reader) (conn.IPacket, error)
}

// IPacketCodec 数据包编解码器
// 负责数据包的解析和编出
type IPacketCodec interface {
	PacketEncoder
	PacketDecoder
}

// DefaultPacketCodec 数据报默认的编解码器
type DefaultPacketCodec struct{}

// Encode 编码
func (d *DefaultPacketCodec) Encode(packet conn.IPacket) ([]byte, error) {
	idBuff := make([]byte, 4)
	binary.BigEndian.PutUint32(idBuff, packet.GetID())

	sizeBuff := make([]byte, 4)
	binary.BigEndian.PutUint32(sizeBuff, uint32(len(packet.GetData())))

	idBuff = append(idBuff, sizeBuff...)
	idBuff = append(idBuff, packet.GetData()...)

	return idBuff, nil
}

// Decode 解码
func (d *DefaultPacketCodec) Decode(reader io.Reader) (conn.IPacket, error) {
	idBuff := make([]byte, 4)
	if _, err := io.ReadFull(reader, idBuff); err != nil {
		return nil, err
	}
	id := binary.BigEndian.Uint32(idBuff)
	sizeBuff := make([]byte, 4)
	if _, err := io.ReadFull(reader, sizeBuff); err != nil {
		return nil, err
	}
	size := binary.BigEndian.Uint32(sizeBuff)
	data := make([]byte, size)
	if _, err := io.ReadFull(reader, data); err != nil {
		return nil, err
	}

	p := new(conn.DefaultPacket)
	p.SetID(id)
	p.SetData(data)

	return p, nil
}
