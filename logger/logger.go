package logger

import (
	"fmt"
	"io"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 15:15
* @version: 1.0
* @description:
*********************************************************/

type Logger struct {
	writer io.Writer
}

func (l *Logger) Log(a ...any) {
}

func NewLogger(writer io.Writer) *Logger {
	return &Logger{writer: writer}
}

func Log(a ...any) {
	fmt.Println(a)
}
