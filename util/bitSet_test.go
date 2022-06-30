package util

import (
	"fmt"
	"testing"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/30 0030 12:02
* @version: 1.0
* @description:
*********************************************************/

func Test_BitSet(t *testing.T) {
	var i8 uint8
	var i8_2 uint8 = 255
	i8 = i8_2 << 4
	fmt.Println(i8) // f0
	i8_2 <<= 4
	fmt.Println(i8_2) // f0
	a := 6888
	var j uint64 = (1 << a) - 1
	var k uint64 = (1 << uint64(62)) - 1
	fmt.Printf("%b\n", k)
	fmt.Printf("%b%v\n", j, j == 0xffffffffffffffff)
	// 若范围为10
	//
	// 111111...00
	// 000000...01
	// 000000...10
	fmt.Printf("%08b", ^9)

}
