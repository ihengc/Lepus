package util

/********************************************************
* @author: Ihc
* @date: 2022/6/30 0030 16:52
* @version: 1.0
* @description:
*********************************************************/

// IList 有序集合接口
type IList[E any] interface {
	Size() int
	IsEmpty() bool
	Contains(e E) bool
	ContainsAll(es []E) bool
	Add(e E) bool
	AddAll(es []E) bool
	Remove(e E) bool
	RemoveByIndex(index int) E
	RemoveAll(es []E) bool
	Get(index int) E
	Set(index int, e E) E
	AddByIndex(index int, e E)
	IndexOf(e E) int
	LastIndexOf(e E) int
	SubList(fromIndex, toIndex int) IList[E]
}
