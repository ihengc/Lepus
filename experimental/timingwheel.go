package experimental

import (
	"sync"
	"time"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/23 22:31
 * @description:
 ***************************************************************/

// TimerTask 定时任务
type TimerTask struct {
	// expiration 过期时间
	expiration time.Duration
}

// TimerTaskNode 定时器任务列表节点
type TimerTaskNode struct {
	prev *TimerTaskNode
	next *TimerTaskNode
	task *TimerTask
}

// TimerTaskList 定时器任务列表
// 用于存放定时任务，定时任务存放
// 在 TimerTaskNode 中，其中节
// 点按过期时间从小到大排列
type TimerTaskList struct {
	// root 根节点，根节点不存放定时任务
	root *TimerTaskNode
}

// TimingWheel 时间轮定时器
type TimingWheel struct {
	// lock 使定时器能用于多线程环境中
	lock sync.RWMutex
	// tickMs 时间轮精度
	tickMs time.Duration
	// wheelSize 时间轮槽数
	wheelSize int
	// buckets 时间槽
	buckets []*TimerTaskList
}
