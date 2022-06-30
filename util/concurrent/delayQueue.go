package concurrent

import (
	"Lepus/util"
	"sync"
	"time"
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

type IDelayQueueElement interface {
	Compare(element IDelayQueueElement) int
}

// DelayQueue 延迟队列
type DelayQueue[E IDelayQueueElement] struct {
	lock  sync.Mutex
	queue util.PriorityQueue[E]
}

// Add 插入指定元素到延迟队列中
// 不允许插入值为nil的值。成功插入返回true；否则返回false
func (d *DelayQueue[E]) Add(e E) bool {
	return d.Offer(e)
}

// Offer 插入指定元素到延迟队列中
// 不允许插入值为nil的值。成功插入返回true；否则返回false
func (d *DelayQueue[E]) Offer(e E) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.queue.Offer(e)
}

// Put 将指定的元素插入此延迟队列。由于队列是无界的，所以此方法永远不会阻塞。
func (d *DelayQueue[E]) Put(e E) {

}

// Poll 检索并删除此队列的头，如果此队列没有延迟过期的元素，则返回null
func (d *DelayQueue[E]) Poll() E {
	return nil
}

// PollWait 检索并删除此队列的标头，如有必要，请等待，直到此队列上具有过期延迟的元素可用
// 或者指定的等待时间过期。timeout表示至多等待此时间
func (d *DelayQueue[E]) PollWait(timeout time.Duration) {

}

// Peek 检索但不删除此队列的头，如果此队列为空，则返回null。与PollWait不同，如果队列中没有可
// 用的过期元素，则此方法返回下一个过期的元素（如果存在）。
func (d *DelayQueue[E]) Peek() E {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.queue.Peek()
}

// Take 检索并删除此队列的头，如有必要，等待此队列上具有过期延迟的元素可用。
func (d *DelayQueue[E]) Take() E {
	return nil
}

// Clear 以原子方式从此延迟队列中删除所有元素。此调用返回后，队列将为空。不等待延迟未过期的元素；
// 它们只是从队列中丢弃。
func (d *DelayQueue[E]) Clear() {

}

// Remove 从此队列中删除指定元素的单个实例（如果存在），无论其是否已过期。
func (d *DelayQueue[E]) Remove() bool {
	return true
}

// Size 返回队列中的元素个数
func (d *DelayQueue[E]) Size() int {
	return d.queue.Size()
}
