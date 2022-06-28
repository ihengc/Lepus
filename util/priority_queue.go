package util

/********************************************************
* @author: Ihc
* @date: 2022/6/28 0028 09:36
* @version: 1.0
* @description:
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

type Option func(*PriorityQueue)

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

type PriorityQueue struct {
	size       int
	comparator Comparator
	queue      []interface{}
}

// grow 调整底层数据大小
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

// Add 将元素e插入到队列中
func (p *PriorityQueue) Add(e interface{}) (bool, error) {
	return p.Offer(e)
}

// Offer 将元素e插入到队列中
// 不允许插入为nil的元素
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

// siftUp 从指定位置处调整元素，使其满足堆的性质
func (p *PriorityQueue) siftUp(k int, x interface{}) {
	for k > 0 {
		parentIndex := (k - 1) >> 1
		parent := p.queue[parentIndex]
		if p.comparator.Compare(parent, x) <= 0 {
			break
		}
		p.queue[k] = parent
		k = parentIndex
	}
	p.queue[k] = x
}

// siftDown 保持位置k及以下的节点堆的特性
func (p *PriorityQueue) siftDown(k int, x interface{}) {
	n := p.size
	half := n >> 1
	for k < half {
		childIndex := (k << 1) + 1
		child := p.queue[childIndex]
		rightIndex := childIndex + 1
		if rightIndex < n && p.comparator.Compare(child, p.queue[rightIndex]) > 0 {
			child = p.queue[childIndex]
		}
		if p.comparator.Compare(x, child) <= 0 {
			break
		}
		p.queue[k] = child
		k = childIndex
	}
	p.queue[k] = x
}

// indexOf 返回元素e在底层数组中的索引
// 若元素e不在数组中，则返回-1
func (p *PriorityQueue) indexOf(e interface{}) int {
	if e != nil {
		n := p.size
		for i := 0; i < n; i++ {
			if p.comparator.Compare(p.queue[i], e) == 0 {
				return i
			}
		}
	}
	return -1
}

// removeAt 移除指定索引处的元素
// 返回被移除的元素
func (p *PriorityQueue) removeAt(i int) interface{} {
	ns := p.size - 1
	if ns == i {
		p.queue[i] = nil
	} else {

	}
	return nil
}

// Remove 从队列中移除指定元素的单个实例(若元素存在)
func (p *PriorityQueue) Remove(e interface{}) bool {
	i := p.indexOf(e)
	if i == -1 {
		return false
	} else {
		p.removeAt(i)
		return true
	}
}

// Contains 若此队列包含元素e，则返回true；否则返回false
func (p *PriorityQueue) Contains(e interface{}) bool {
	return p.indexOf(e) >= 0
}

// Peek 检索但不删除此队列的头部，如果此队列为空，则返回nil
func (p *PriorityQueue) Peek() interface{} {
	return p.queue[0]
}

// Poll 检索并删除此队列的头部，如果此队列为空，则返回nil
func (p *PriorityQueue) Poll() interface{} {
	return nil
}

// Clear 清空队列
// 保持底层数组元素个数和容量不变
// 不更换底层数组
func (p *PriorityQueue) Clear() {
	for i := 0; i < cap(p.queue); i++ {
		p.queue[i] = nil
	}
}

// Size 返回队列中的元素个数
func (p *PriorityQueue) Size() int {
	return p.size
}

// SetComparator 设置比较器
func SetComparator(comparator Comparator) Option {
	if comparator == nil {
		panic(NilPointErr)
	}
	return func(queue *PriorityQueue) {
		queue.comparator = comparator
	}
}

// NewPriorityQueue 创建优先级队列
func NewPriorityQueue(initialCapacity int, options ...Option) *PriorityQueue {
	p := new(PriorityQueue)
	for _, opt := range options {
		opt(p)
	}
	if initialCapacity <= 0 {
		p.queue = make([]interface{}, 11, 11)
	} else {
		p.queue = make([]interface{}, initialCapacity, initialCapacity)
	}
	return p
}
