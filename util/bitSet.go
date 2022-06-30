package util

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/29 21:43
 * @description: BitSet
 ***************************************************************/

// BitSet为存放bit的容器，会涉及到位的逻辑运算和掩码相关概念
// 常用的位的逻辑运算：
// 	1.与运算
// 		与一个值为假的数做与运算，其值必为假
//		与一个值为真的数做与运算，其值保持原样
//		需要将一个值置为假时，可以用假值与其做与运算
// 	2.或运算
//		与一个值为真的数值做或运算，其值必为真
//		需要将一个值置为真时，可以用真值与其做或运算
// 	3.异或运算
// 		与一个值为真的数做异或运算，其值为该数值的相反
//		需要将一个值反转时，可以用真值与其做异或运算
// 		与一个值位假的数做异或运算，真值被反转，假值保持不变

// IBitSet 定义BitSet需要实现的接口
type IBitSet interface {
	// Set 指定索引处的位设置为true。
	// 若设置成功返回true；否则返回false
	Set(bitIndex int) bool

	// SetValue 指定索引处的位设置为给定的值。
	// 若设置成功则返回true；否则返回false
	SetValue(bitIndex int, value bool) bool

	// SetRange 设置fromIndex到toIndex范围内的位设置为true。
	// 若设置成功返回true；否则返回false
	SetRange(fromIndex, toIndex int) bool

	// SetValueRange 设置fromIndex到toIndex范围内的位为指定的值（0或1）。
	// 若设置成功返回true；否则返回false
	SetValueRange(fromIndex, toIndex int, value bool) bool

	// NextSetBit
	// 返回第一个被设置为true的位的索引
	NextSetBit(fromIndex int) int

	// ClearAll 设置所有的位为false
	ClearAll()

	// Clear 指定索引处的位设置为false
	// 若设置成功返回true；否则返回false
	Clear(bitIndex int) bool

	// ClearRange 将fromIndex到toIndex范围内的位设置为false。
	// 若设置成功返回true；否则返回false
	ClearRange(fromIndex, toIndex int) bool

	// NextClearBit
	// 返回第一个被设置为false的位的索引
	NextClearBit(fromIndex int) int

	// Get 获取指定索引处位的值
	Get(bitIndex int) (bool, error)

	// GetRange 获取fromIndex到toIndex范围内位的值
	// 若指定范围合法，返回BitSet；否则返回nil
	GetRange(fromIndex, toIndex int) IBitSet

	// Flip 将指定索引处的位设置为其补码
	Flip(bitIndex int)

	// FlipRange 将指定范围索引内的位设置为其补码
	FlipRange(fromIndex, toIndex int) bool

	// IsEmpty 若BitSet中无位被设置为true则返回true
	IsEmpty() bool

	// Len 返回BitSet的逻辑大小
	Len() int

	// And 与目标BItSet执行逻辑与运算
	And(set IBitSet)

	// AndNot

	// Or
	// Xor
}

const (
	wordMask = 0xffffffffffffffff
)

// BitSet bit容器
type BitSet struct {
	// words 用来存放位，为int64数组，数组中每
	// 一个元素占64位
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

// Set 指定索引处的位设置为true
func (b *BitSet) Set(bitIndex int) bool {
	if bitIndex < 0 {
		return false
	}
	wordIndex := b.wordIndex(bitIndex)
	b.expandTo(wordIndex)

	b.words[wordIndex] |= 1 << b.bitIndexOffset(bitIndex)
	return true
}

// SetValue 指定索引处的位置设置为value值（0或1）
func (b *BitSet) SetValue(bitIndex int, value bool) bool {
	if bitIndex < 0 {
		return false
	}
	if value {
		b.Set(bitIndex)
	} else {
		b.Clear(bitIndex)
	}
	return true
}

// Clear 将指定索引处位的值设置为false
// 与一个假值做与运算时，其得到的值必定为假；所以可以采用与运算
// 但是我们需要考虑的除了要处理的位其他位是需要保持原样的。在与
// 运算中只有与一个真值做与运算得到的值才会是其本身。
// 所以我们需要类似 01111 1111... 的二进制数。我们只需要将1向
// 左进行移位，然后取反即可
func (b *BitSet) Clear(bitIndex int) bool {
	if bitIndex < 0 {
		return false
	}
	// wordIndex 表示当前bitIndex应该在第几个元素中
	wordIndex := b.wordIndex(bitIndex)
	if wordIndex >= b.wordsInUse {
		return false
	}
	b.words[wordIndex] &= ^(1 << b.bitIndexOffset(bitIndex))
	// 将一位置为了false，需要重新计算BitSet逻辑大小
	b.recalculateWordsInUse()
	return true
}

// ClearRange 将fromIndex到toIndex范围内的位设置为false。
// 我们需要将一个int64位中的某一段位置为false。假设范围的起
// 始位置和结束位置分别为startIndex和endIndex，我们需要在
// startIndex到endIndex之间所有位的值为false，其他位为false
// 如下:
// 		1111-0000-0011
// = 	1111-0000-0000 +
// 		0000-0000-0011
// 我们只需要将全为true的int64位数左移到startIndex，
//
func (b *BitSet) ClearRange(fromIndex, toIndex int) bool {
	if !b.checkRange(fromIndex, toIndex) {
		return false
	}
	startWordIndex := b.wordIndex(fromIndex)
	endWordIndex := b.wordIndex(toIndex)
	// startWordIndex等于endWordIndex说明区间在一个int64内
	if startWordIndex == endWordIndex {
		startOffset := b.bitIndexOffset(fromIndex)
		endOffset := b.bitIndexOffset(toIndex)
		b.words[startWordIndex] &= wordMask & (1 << endOffset) & (1<<startOffset - 1)
	} else {
		//
		for i := startWordIndex; i < endWordIndex; i++ {
			b.words[i] = 0
		}
	}
	return true
}

func NewBitSet(nBits int) *BitSet {
	b := new(BitSet)
	b.initWords(nBits)
	b.sizeIsSticky = false
	return b
}
