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
	ID   uint32
	Data []byte
}

func (d *DefaultPacket) GetID() uint32 {
	//TODO implement me
	panic("implement me")
}

func (d *DefaultPacket) GetData() []byte {
	//TODO implement me
	panic("implement me")
}

func (d *DefaultPacket) SetID(id uint32) {
	//TODO implement me
	panic("implement me")
}

func (d *DefaultPacket) SetData(data []byte) {
	//TODO implement me
	panic("implement me")
}
