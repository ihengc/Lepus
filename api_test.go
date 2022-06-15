package Lepus

import (
	"testing"
	"time"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/15 0015 14:53
* @version: 1.0
* @description:
*********************************************************/

var lepus *Lepus

func setup() {
	lepus = &Lepus{}
	lepus.Stmt = &Statement{Fields: make([]*Field, 0)}
	lepus.sqlBuilder = &SQLBuilder{}
}

type Player struct {
	ID       uint
	Name     string
	Age      int
	Birthday time.Time
}

func TestLepus_Create(t *testing.T) {
	setup()
	player := &Player{ID: 1, Name: "Lepus", Age: 20, Birthday: time.Now()}
	lepus.Create(player)
}
