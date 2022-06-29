package util

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/29 21:43
 * @description: BitSet
 ***************************************************************/

// BitSet为存放bit的容器，会涉及到位的逻辑运算和掩码相关概念
// 常用的位的逻辑运算：
// 1.与运算
// 2.或运算
// 3.非运算
// 4.异或运算
// 与一个值为真的数做异或运算，则返回该数的相反。
// 与一个值位假的数做异或运算，真值被反转，假值保持不变
// 关于掩码：
// 关于左移：
// 在Python中左移操作是没有溢出的情况的，而在Golang中左移会
// 发生溢出，当左移后的数大于当前数的数据类型最大值时，会报错
// 如下代码会有问题：
// int8(1) << 7
// 第8位为符号位，所以最多只能移动到第7位。

// IBitSet 定义BitSet需要实现的接口
type IBitSet interface {
}

const (
	wordMask = 0xffffffffffffffff
)

// BitSet bit容器
type BitSet struct {
	words []int64
	// wordsInUse words的逻辑大小
	wordsInUse int
	// sizeIsSticky
	sizeIsSticky bool
}

// recalculateWordsInUse
func (b *BitSet) recalculateWordsInUse() {
	var i int
	for i = b.wordsInUse - 1; i >= 0; i-- {
		if b.words[i] != 0 {
			break
		}
	}
	b.wordsInUse = i + 1
}

// wordIndex 将指定bit的位置转换位int64数组的索引
func (b *BitSet) wordIndex(nBits int) int {
	return nBits >> 6
}

// ensureCapacity
func (b *BitSet) ensureCapacity(wordsRequired int) {
	wordsCapacity := cap(b.words)
	if wordsCapacity < wordsRequired {
		var requestSize int
		if 2*wordsCapacity > wordsRequired {
			requestSize = 2 * wordsCapacity
		} else {
			requestSize = wordsRequired
		}
		newWords := make([]int64, requestSize, requestSize)
		copy(newWords, b.words)
		b.words = newWords
	}
}

// expandTo
func (b *BitSet) expandTo(wordIndex int) {
	wordsRequired := wordIndex + 1
	if b.wordsInUse < wordsRequired {
		b.ensureCapacity(wordsRequired)
		b.wordsInUse = wordsRequired
	}
}

// checkRange 检查fromIndex ...toIndex 是否合理
func (b *BitSet) checkRange(fromIndex, toIndex int) bool {
	if fromIndex < 0 || toIndex < 0 || fromIndex < toIndex {
		return false
	}
	return true
}

// initWords
func (b *BitSet) initWords(nBits int) {
	b.words = make([]int64, b.wordIndex(nBits), b.wordIndex(nBits))
}

// bitIndexOffset
func (b *BitSet) bitIndexOffset(bitIndex int) int {
	return bitIndex % 64
}

// Flip 翻转指定位置处的比特位
// 若翻转成功，则返回true；否则返回false
func (b *BitSet) Flip(bitIndex int) bool {
	if bitIndex < 0 {
		return false
	}
	wordIndex := b.wordIndex(bitIndex)
	b.expandTo(wordIndex)

	b.words[wordIndex] ^= 1 << b.bitIndexOffset(bitIndex)
	return true
}

// FlipRange 翻转
func (b *BitSet) FlipRange(fromIndex, toIndex int) {
	startWordIndex := b.wordIndex(fromIndex)
	endWordIndex := b.wordIndex(toIndex)

	if startWordIndex == endWordIndex {
		b.words[startWordIndex] ^= 1
	}
}

func NewBitSet(nBits int) *BitSet {
	b := new(BitSet)
	b.initWords(nBits)
	b.sizeIsSticky = false
	return b
}
