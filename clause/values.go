package clause

/********************************************************
* @author: Ihc
* @date: 2022/6/15 0015 16:37
* @version: 1.0
* @description:
*********************************************************/

type Values struct {
	Values [][]interface{}
}

func (v Values) Name() string {
	return "VALUES"
}

func (v Values) Build(clauseBuilder IClauseBuilder) {
	for idx, vList := range v.Values {
		if idx > 0 {
			clauseBuilder.WriteByte(',')
		}
		clauseBuilder.WriteByte('(')
		for idx, value := range vList {
			if idx > 0 {
				clauseBuilder.WriteByte(',')
			}
			clauseBuilder.WriteQuoted(value)
		}
		clauseBuilder.WriteByte(')')
	}
}
