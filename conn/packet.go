package conn

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/17 20:26
 * @description:
 ***************************************************************/

// IPacket 表示一个数据包
type IPacket interface {
	GetID() int
	GetData() []byte
	SetID(int)
	SetData([]byte)
}
