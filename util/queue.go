package util

/********************************************************
* @author: Ihc
* @date: 2022/6/30 0030 17:01
* @version: 1.0
* @description:
*********************************************************/

// IQueue 队列接口
type IQueue[E any] interface {
	// Add 将指定元素插入到队列
	// 若队列容量充足返回true；否则返回false
	Add(e E) bool

	// Element 检索但不删除队列的头部
	// 若队列为空返回nil；否则返回队列头部的复制
	Element() E

	// Offer 如果可以在不违反容量限制的情况下立即将指定的元素插入到队列中
	Offer(e E) bool

	// Poll 检索并删除该队列的头部，如果该队列为空，则返回null
	Poll() E

	// Remove 检索并删除该队列的头部。进入翻译页面
	Remove() E
}
