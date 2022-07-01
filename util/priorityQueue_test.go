package util

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/7/1 0001 15:25
* @version: 1.0
* @description:
*********************************************************/

var p *PriorityQueue[pElement]

type pElement struct {
	value int
}

type pElementComparator struct{}

func (p *pElementComparator) Compare(e1, e2 interface{}) int {
	pe1 := e1.(*pElement)
	pe2 := e2.(*pElement)

	if pe1.value > pe2.value {
		return 1
	}
	if pe1.value < pe2.value {
		return -1
	}
	return 0
}

func newPriorityQueue() {
	var err error
	p, err = NewPriorityQueue[pElement](10, &pElementComparator{})
	if err != nil {
		panic(err)
	}
}

func init() {
	newPriorityQueue()
}

func checkPriorityQueueSize(queue *PriorityQueue[pElement], size int) bool {
	return queue.Size() == size
}

func TestPriorityQueue_Add(t *testing.T) {
	e := &pElement{value: 10}
	p.Add(e)
	if checkPriorityQueueSize(p, 1) {

	}

}
