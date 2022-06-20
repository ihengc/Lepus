package logger

import (
	"fmt"
	"io"
	"sync"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 15:15
* @version: 1.0
* @description:
*********************************************************/

type Logger struct {
	out    io.Writer
	lock   sync.Mutex
	buffer []byte
	prefix string
}

func Log(a ...any) {
	fmt.Println(a)
}
