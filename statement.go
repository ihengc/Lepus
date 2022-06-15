package Lepus

import "strings"

/********************************************************
* @author: Ihc
* @date: 2022/6/15 0015 14:36
* @version: 1.0
* @description:
*********************************************************/

type Statement struct {
	Dest interface{}
	SQL  strings.Builder
}
