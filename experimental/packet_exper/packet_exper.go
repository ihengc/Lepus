package packet_exper

import "io"

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/21 21:46
 * @description:
 ***************************************************************/

// 用户自定义数据包的编解码器的好处与坏处。
// 用户自定义数据包的编解码可以定制数据包的数据结构
// 但也较为繁琐

type PacketType byte

const (
	// SYN  ACK  用于UDP协议中，UDP不是面向廉连接的协议
	SYN PacketType = iota + 1
	ACK
	Heartbeats
	Data
)

// IPacketEncoder 数据包编码接口
// IPacketDecoder 数据包解码接口

// IPacketCodec 数据包的编解码接口
type IPacketCodec interface {
	// Encode 将数据包对象编码为二进制，编码应该注意与发送端协商使用何种字节序
	// 通常TCP/IP中规定使用大端字节序
	Encode(IPacket) ([]byte, error)

	// Decode 将二进制解析为数据包对象
	Decode(io.Reader) (IPacket, error)
}

// IPacket 数据包接口
type IPacket interface {
	GetServerName() string     // 获取所属服务的名称
	GetData() []byte           // 获取数据包正文
	GetPacketType() PacketType // 获取数据包类型
	SetData([]byte)            // 设置数据包正文
}

// BasePacket 数据包
type BasePacket struct {
	Data       []byte     // 数据正文
	ServerName string     // 若当前数据包处于分布式系统中，需要指明这个包应该在何种服务上被处理
	PacketType PacketType // 数据包类型
}

// TCPPacket 使用TCP协议发送的数据包
// 使用TCP协议传输数据，会涉及到分包的问题
type TCPPacket struct {
	BasePacket
	size uint32 // 指明数据包的大小
}

// UDPPacket 使用UDP协议发送的数据包
// UDP协议的数据包
type UDPPacket struct {
	BasePacket
	Sequence uint32 // 序列号，UDP数据报的到达是无序的，通过该字段来对到达的数据报进行排序
}
