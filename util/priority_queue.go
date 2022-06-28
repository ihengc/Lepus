package util

/********************************************************
* @author: Ihc
* @date: 2022/6/28 0028 09:36
* @version: 1.0
* @description: 优先级队列
*********************************************************/
import (
	"errors"
)

var (
	NilPointErr = errors.New("nil point error")
)

// Comparator 比较接口
type Comparator interface {
	// Compare 比较e1与e2的大小
	Compare(e1, e2 interface{}) int
}

type IPriorityQueue interface {
	// Add 将指定元素插入到队列中
	Add(e interface{}) (bool, error)
	// Offer 将指定元素插入到队列中
	Offer(e interface{}) (bool, error)
	// Remove 从队列中移除指定元素的单个实例(若元素存在)
	Remove(e interface{}) bool
	// Clear 清空队列
	Clear()
	// Size 返回队列中的元素个数
	Size() int
	// Poll 检索并删除此队列的头部，如果此队列为空，则返回nil
	Poll() interface{}
	// Peek 检索但不删除此队列的头部，如果此队列为空，则返回nil
	Peek() interface{}
	// Contains 如果此队列包含指定元素，则返回 true
	Contains(e interface{}) bool
}

// PriorityQueue 优先级队列
type PriorityQueue struct {
	// size 存放优先级队列中的元素大小
	size int
	// comparator 元素比较器
	comparator Comparator
	// queue 存放元素的数组
	queue []interface{}
}

// grow 数组扩容
// 底层的数组需要定义初始容量，扩容需要有对应
// 的扩容机制。当数组容量小于64时，每次扩容2
// 个容量；若数组容量大于等于64时，每次扩容原
// 来容量的1.5倍
func (p *PriorityQueue) grow() {
	var (
		oldCapacity int
		newCapacity int
		newQueue    []interface{}
	)
	oldCapacity = cap(p.queue)
	if oldCapacity < 64 {
		newCapacity = oldCapacity + 2
	} else {
		newCapacity = oldCapacity + oldCapacity>>1
	}
	newQueue = make([]interface{}, newCapacity, newCapacity)
	copy(newQueue, p.queue)
	p.queue = newQueue
}

// Offer 将指定元素插入到队列中
func (p *PriorityQueue) Offer(e interface{}) (bool, error) {
	if e == nil {
		return false, NilPointErr
	}
	if p.size >= cap(p.queue) {
		p.grow()
	}
	p.siftUp(p.size, e)
	p.size++
	return true, nil
}

// siftUp 向上调整树的结构，使其满足堆的性质
// 对于插入一个元素到堆中我们关心的是，新元素
// 的加入能否让数组继续满足堆的性质。如何让数
// 组继续满足堆的性质？因为树是一个递归定义的
// 结构，所以我们只用关系被递归的单元是否满足
// 堆的性质即可，也就是关心插入元素与其父节点
// 值的大小。
// (小根堆)
// 若插入元素的值比其父节点的值要小，则不满足
// 堆的性质，需要将其父节点的值移动到插入元素
// 的位置。反之则说明已经满足了堆的性质。当整
// 个树已经满足了堆的性质，也就是说明我们找到
// 了新增元素正确地插入位置，我们记录插入的位
// 置，然后在该位置放入新增元素即可。
func (p *PriorityQueue) siftUp(k int, x interface{}) {
	for k > 0 {
		parentIndex := (k - 1) >> 1
		parent := p.queue[parentIndex]
		if p.comparator.Compare(x, parent) >= 0 {
			break
		}
		p.queue[k] = parent
		k = parentIndex
	}
	p.queue[k] = x
}
