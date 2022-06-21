package codec

import (
	"Lepus/conn"
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
