package conn

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/17 20:26
 * @description:
 ***************************************************************/

// IPacket 表示一个数据包
type IPacket interface {
	GetID() uint32
	GetData() []byte
	SetID(uint32)
	SetData([]byte)
}

// DefaultPacket 数据包
type DefaultPacket struct {
	id   uint32
	data []byte
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
