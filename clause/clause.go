package clause

/********************************************************
* @author: Ihc
* @date: 2022/6/15 0015 16:35
* @version: 1.0
* @description:
*********************************************************/

type IClause interface {
	Build(builder IClauseBuilder)
}

type IClauseBuilder interface {
	WriteByte(byte)
	WriteString(string)
	WriteQuoted(interface{})
}
