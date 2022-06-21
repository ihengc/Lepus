package codec

import "Lepus/conn"

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/20 22:45
 * @description:
 ***************************************************************/

type PacketEncoder interface {
	Encode(packet *conn.Packet) ([]byte, error)
}
