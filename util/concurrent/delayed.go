package concurrent

import "time"

/********************************************************
* @author: Ihc
* @date: 2022/6/29 0029 15:59
* @version: 1.0
* @description:
*********************************************************/

// IDelayed 用于标记在给定延迟后应该执行的对象
type IDelayed interface {
	// GetDelay 返回在给定时间单位内与该对象关联的剩余延迟时间。
	GetDelay(duration time.Duration) int64
}
