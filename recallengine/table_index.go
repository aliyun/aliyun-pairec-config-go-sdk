package recallengine

import (
	"fmt"
	"sync/atomic"
)

const (
	// DocIndexMask, 行数应小于 16777215 (0x00ffffff)
	DocIndexMask uint32 = 0x00ffffff

	// DocDeleteMask 用于标记一行已被删除
	DocDeleteMask uint32 = 0x80000000
)

// TableIndex 维护用于 Record 的数据位置。
// 它使用一个整数切片来保存数据位置。
// 当数据写入 Record 时，其在 IColumn 中的数据位置是固定的。
// 当排序或删除数据时，IColumn 中的数据不受影响，但我们会调整 TableIndex。
//
// 示例：
// 假设记录中有 5 个元素，初始索引为：
//
//	0 -> 0
//	1 -> 1
//	2 -> 2
//	3 -> 3
//	4 -> 4
//
// 左边的值是 indexes 切片的下标，右边的值是数据在记录中的实际位置。
//
// 当我们对数据进行排序后，索引可能会变成这样：
//
//	0 -> 2
//	1 -> 0
//	2 -> 3
//	3 -> 1
//	4 -> 4
//
// 因此，在输出记录元素时，IColumn 中位置为 2 的元素将第一个被输出。
//
// 当数据被删除时，我们会通过位操作 indexes[index] = indexes[index] | DocDeleteMask 来标记。
type TableIndex struct {
	// 使用 atomic 来安全地读写 rowCount
	rowCount int32
	indexes  []uint32
}

// NewTableIndex 创建并初始化一个新的 TableIndex。
// count 是初始行数。如果 count 无效，则返回错误。
func NewTableIndex(count int) *TableIndex {
	if count > int(DocIndexMask) {
		count = int(DocIndexMask)
	}
	if count < 0 {
		count = 0
	}

	ti := &TableIndex{
		rowCount: int32(count),
		indexes:  make([]uint32, count),
	}

	for i := 0; i < count; i++ {
		ti.indexes[i] = uint32(i) & DocIndexMask
	}
	return ti
}

// RowCount 返回当前的行数。此操作是并发安全的。
func (ti *TableIndex) RowCount() int {
	return int(ti.rowCount)
}

// Indexes 返回内部索引切片的一个副本。
// 返回副本是为了防止外部代码在没有锁的情况下修改内部状态。
func (ti *TableIndex) Indexes() []uint32 {
	// 创建并返回一个副本
	copiedIndexes := make([]uint32, len(ti.indexes))
	copy(copiedIndexes, ti.indexes)
	return copiedIndexes
}

// GetIndex 获取给定行在 IColumn 中的真实位置。
func (ti *TableIndex) GetIndex(row int) int {
	return int(ti.indexes[row] & DocIndexMask)
}
func (ti *TableIndex) SafeGet(row int) int {
	return int(ti.indexes[row] & DocIndexMask)
}

// MarkRemoved 将指定索引处的行标记为已删除。
func (ti *TableIndex) MarkRemoved(index int) {
	ti.indexes[index] |= DocDeleteMask
}

// IsRemoved 检查给定行是否已被标记为删除。
func (ti *TableIndex) IsRemoved(row int) bool {
	return (ti.indexes[row] & DocDeleteMask) != 0
}

// Rebuild 使用一个新的索引切片替换原有的索引。
func (ti *TableIndex) Rebuild(rowCount int, indexes []uint32) {
	ti.indexes = indexes
	atomic.StoreInt32(&ti.rowCount, int32(rowCount))
}

// Truncate 截断行数，保留前面的 n 行。
func (ti *TableIndex) Truncate(n int) error {
	if n <= 0 {
		return fmt.Errorf("illegal table index truncate n: %d", n)
	}
	// 委托给更通用的 TruncateFrom
	return ti.TruncateFrom(0, n)
}

// TruncateFrom 从 from 位置开始截断，保留 n 行。
func (ti *TableIndex) TruncateFrom(from, n int) error {

	currentCount := int(atomic.LoadInt32(&ti.rowCount))

	if from >= currentCount {
		// 如果 'from' 超出范围，结果是一个空索引
		ti.indexes = ti.indexes[:0]
		atomic.StoreInt32(&ti.rowCount, 0)
		return nil
	}

	if n < 0 || from < 0 {
		return fmt.Errorf("illegal table index truncate from: %d, n: %d", from, n)
	}

	// 如果 n 超出从 from 开始的剩余元素数量，则调整 n
	if n > currentCount-from {
		n = currentCount - from
	}

	// 复制所需的段到切片的开头
	copy(ti.indexes, ti.indexes[from:from+n])
	// 重新切片以更新其长度
	ti.indexes = ti.indexes[:n]

	atomic.StoreInt32(&ti.rowCount, int32(n))
	return nil
}

// EndRemove 压缩切片，移除所有被标记为删除的行。
// 这不是真正的删除，而是将未删除的行移动到前面并更新行数。
func (ti *TableIndex) EndRemove() {

	currentCount := int(atomic.LoadInt32(&ti.rowCount))
	writeIndex := 0
	// 使用经典的 "in-place" 过滤算法
	for readIndex := 0; readIndex < currentCount; readIndex++ {
		// 如果行未被删除
		if (ti.indexes[readIndex] & DocDeleteMask) == 0 {
			// 如果写入位置和读取位置不同，则进行移动
			if writeIndex != readIndex {
				ti.indexes[writeIndex] = ti.indexes[readIndex]
			}
			writeIndex++
		}
	}

	// 重新切片以反映新的大小
	ti.indexes = ti.indexes[:writeIndex]
	atomic.StoreInt32(&ti.rowCount, int32(writeIndex))
}

// Get 返回索引 i 处的原始值（包括删除标记）。
func (ti *TableIndex) Get(i int) uint32 {
	return ti.indexes[i]
}
