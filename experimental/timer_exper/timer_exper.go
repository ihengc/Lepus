package timer_exper

import "time"

/********************************************************
* @author: Ihc
* @date: 2022/6/23 0023 09:56
* @version: 1.0
* @description:
*********************************************************/

// 定时器如何存放定时任务？
// 定时器如何同一时刻执行多个定时任务？
// 定时器如何重复执行定时任务？

// 使用双向循环链表来存放定时任务，双向循环链表数据结构：
// 1）根节点，哨兵节点，不存放数据
// 2) 过期时间，时间轮中的槽位

// ITimer 定时器接口
type ITimer interface {
	// Add 添加定时任务
	Add(*TimerTask)
}

// ITimerTaskList 存放定时任务链表接口
type ITimerTaskList interface {
	// Add 添加定时任务记得
	Add(entry *TimerTaskEntry)
	// Remove 删除定时任务节点
	Remove(entry *TimerTaskEntry)
	// GetDelay 获取延迟时间
	GetDelay() time.Duration
}

// ITimerTaskEntry 定时任务链表节点接口
type ITimerTaskEntry interface {
	// Cancelled 定时任务是否被取消
	Cancelled() bool
	// Remove 重定时任务链表中删除此节点
	Remove()
	// Compare 与其他时定时任务节点比较定时任务的过期时间
	Compare(ITimerTaskEntry)
}

// ITimerTask 定时器任务接口
type ITimerTask interface {
	// Cancel 取消该任务
	Cancel()
	// GetTimerTaskEntry 获取与定时任务绑定的定时任务列表节点
	GetTimerTaskEntry() ITimerTaskEntry
	// SetTimerTaskEntry 将定时任务与定时任务列表节点绑定
	SetTimerTaskEntry(ITimerTaskEntry)
}

// TimerTaskList 定时任务列表
type TimerTaskList struct {
	// root 根节点
	root *TimerTaskEntry
}

// newTimerTaskList 创建时间任务列表
func newTimerTaskList() *TimerTaskList {
	t := new(TimerTaskList)

	return t
}

// TimerTaskEntry 定时任务链表节点
type TimerTaskEntry struct {
	// expirationMs 过期时间
	// timerTask 定时任务
}

// TimerTask 定时任务
type TimerTask struct {
	// 任务

}

// TimingWheel 时间轮
type TimingWheel struct {
	// tickMs 时间单位，表示针走一次所经历的时间跨度
	tickMs time.Duration
	// startMs 起始时间
	startMs time.Duration
	// wheelSize 时间轮中槽的数量
	wheelSize int
	// currentTime 当前时间轮经历的时间
	currentTime time.Duration
	// buckets 时间槽，环形队列，里面存放任务列表
	// overflowWheel 上层时间轮

}

// NewTimingWheel 创建时间轮
func NewTimingWheel(wheelSize int, tickMs time.Duration, startMs time.Duration) *TimingWheel {
	t := new(TimingWheel)
	t.wheelSize = wheelSize
	t.tickMs = tickMs
	t.currentTime = startMs - (startMs % tickMs)
	return t
}
