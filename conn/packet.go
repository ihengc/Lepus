package conn

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/17 20:26
 * @description:
 ***************************************************************/

const (
	Heartbeats = iota + 1
)

// IPacket 表示一个数据包
type IPacket interface {
	GetID() uint32
	GetData() []byte
	SetID(uint32)
	SetData([]byte)
	GetServerName() string
	GetPacketType() uint8
}

// DefaultPacket 数据包
type DefaultPacket struct {
	id          uint32
	typ         uint8
	data        []byte
	serviceName string
}

func (d *DefaultPacket) GetID() uint32 {
	return d.id
}

func (d *DefaultPacket) GetData() []byte {
	return d.data
}

func (d *DefaultPacket) SetID(id uint32) {
	d.id = id
}

func (d *DefaultPacket) SetData(data []byte) {
	d.data = data
}

func (d *DefaultPacket) GetServerName() string {
	return "Lepus"
}
