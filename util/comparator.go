package util

/********************************************************
* @author: Ihc
* @date: 2022/6/29 0029 15:51
* @version: 1.0
* @description:
*********************************************************/

// Comparator 比较接口
type Comparator interface {
	// Compare 比较e1与e2的大小
	Compare(e1, e2 interface{}) int
}
