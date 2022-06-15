package clause

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/15 0015 16:46
* @version: 1.0
* @description:
*********************************************************/

type valuesExpr struct {
	SQL strings.Builder
}

func (v *valuesExpr) WriteByte(b byte) {
	v.SQL.WriteByte(b)
}

func (v *valuesExpr) WriteString(s string) {
	v.SQL.WriteString(s)
}

func (v *valuesExpr) WriteQuoted(i interface{}) {
	switch i.(type) {
	case uint8, uint16, uint32, uint64, int8, int16, int32, int64, int, uint:
		v.SQL.WriteString(fmt.Sprintf("%d", i))
	case float32, float64:
		v.SQL.WriteString(fmt.Sprintf("%v", i))
	case string:
		v.SQL.WriteString(fmt.Sprintf("'%s'", i))
	case time.Time:
		t := i.(time.Time)
		v.SQL.WriteString(fmt.Sprintf("'%s'", t.Format("2006-01-02 15:04:05")))
	}
}

func TestValues_Build(t *testing.T) {
	x := valuesExpr{}
	values := &Values{}
	values.Values = [][]interface{}{{1, 2, 3, "4"}, {"5", "6", 7, time.Now(), 12.141, 0.00123456}}
	values.Build(&x)
	t.Log(x.SQL.String())
}
