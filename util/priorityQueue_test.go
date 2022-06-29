package util

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/6/29 0029 11:10
* @version: 1.0
* @description:
*********************************************************/

var queue *PriorityQueue
var cmp *comparable

type comparable struct{}

func (c *comparable) Compare(e1, e2 interface{}) int {
	v1 := e1.(int)
	v2 := e2.(int)
	if v1 == v2 {
		return 0
	} else if v1 < v2 {
		return -1
	}
	return 1
}

func setup() {
	cmp = &comparable{}
	queue = NewPriorityQueue(20, cmp)
}

func TestPriorityQueue(t *testing.T) {
	setup()
	if queue.Size() != 0 {
		t.Fatal("queue size error")
	}

	ret, err := queue.Offer(11)
	if err != nil {
		t.Fatal(err)
	}
	if !ret {
		t.Fatal("queue Offer error")
	}
	if queue.Size() != 1 {
		t.Fatal("queue size error after offer")
	}

	_, err = queue.Offer(12)
	_, err = queue.Offer(1)
	_, err = queue.Offer(2)
	_, err = queue.Offer(12414)
	_, err = queue.Offer(32)
	_, err = queue.Offer(72)
	if err != nil {
		t.Fatal(err)
	}
	if queue.Peek() != 1 {
		t.Fatal("queue peek error")
	}
	if queue.Poll() != 1 {
		t.Fatal("queue poll error")
	}
	if queue.Peek() != 2 {
		t.Fatal("queue error")
	}
}
