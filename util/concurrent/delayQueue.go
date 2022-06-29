package concurrent

import (
	"Lepus/util"
	"sync"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/29 0029 11:21
* @version: 1.0
* @description: 延迟队列
*********************************************************/

// 延迟队列会涉及到线程同步的问题，这里讨论一下可重入锁。
// 首先Golang中的锁是不可重入的，也就是说同一个线程多次
// 获取同一把锁,会造成死锁。如下代码
// func main() {
//		lock := sync.Mutex{}
//		lock.Lock()
//		lock.Lock()
// }
// 当代码中出现可重入锁的需求时，通常是代码设计的问题。解
// 决方法是合并多次获取同一个锁的代码
//
// 可重入锁的实现原理:
//

// IDelayQueue 延迟队列接口
type IDelayQueue interface {
	// Add 插入指定元素到队列中
	Add(e interface{}) bool
	// Offer 插入指定元素到队列中
	Offer(e interface{}) bool
	// Put 插入指定元素到队列中
	Put(e interface{})
	// Take 获取队列头部的元素
	Take() interface{}
}

// DelayQueue 延迟队列
type DelayQueue struct {
	// available 条件变量
	available sync.Cond
	// lock 互斥锁
	lock sync.Mutex
	// queue 优先级队列
	queue *util.PriorityQueue
	// size 队列元素个数
	size int
	// comparator 元素比较器
	comparator util.Comparator
}

func (d *DelayQueue) Add(e interface{}) bool {
	return d.Offer(e)
}

func (d *DelayQueue) Offer(e interface{}) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.queue.Offer(e)
	if d.comparator.Compare(d.queue.Peek(), e) == 0 {
		d.available.Signal()
	}
	return true
}

func (d *DelayQueue) Put(e interface{}) {
	d.Offer(e)
}

func (d *DelayQueue) Peek() interface{} {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.queue.Peek()
}
