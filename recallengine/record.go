package recallengine

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"sort"
	"sync"

	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
)

type Record struct {
	sync.RWMutex
	//size        int
	columnData map[string]IColumn
	tableIndex *TableIndex
	cacheData  map[int]map[string]any
}

func NewRecord(size int) *Record {
	recrod := &Record{
		//size:       size,
		columnData: make(map[string]IColumn),
	}

	tableIndex := NewTableIndex(size)

	recrod.tableIndex = tableIndex

	return recrod
}

func (r *Record) Size() int {
	return r.tableIndex.RowCount()
}
func (r *Record) SetColumn(name string, column IColumn) {
	r.Lock()
	defer r.Unlock()
	r.columnData[name] = column
}
func (r *Record) GetColumn(name string) IColumn {
	r.RLock()
	defer r.RUnlock()
	return r.columnData[name]
}

func (r *Record) Sort(name string, desc bool) {
	r.RLock()
	column, exist := r.columnData[name]
	r.RUnlock()
	if !exist {
		return
	}
	size := r.tableIndex.RowCount()

	sortItems := make(sortItemSlice, 0, size)
	for i := 0; i < size; i++ {
		index := r.tableIndex.SafeGet(i)
		_v, _ := column.Get(index)
		sortItems = append(sortItems, &sortItem{
			index: index,
			value: _v,
		})
	}

	if desc {
		sort.Sort(sort.Reverse(sortItems))
	} else {
		sort.Sort(sortItems)
	}

	indexes := make([]uint32, 0, len(sortItems))
	for _, item := range sortItems {
		indexes = append(indexes, uint32(item.index))
	}
	r.Lock()
	defer r.Unlock()
	r.tableIndex.Rebuild(len(indexes), indexes)
}
func (r *Record) Len() int {
	return r.tableIndex.RowCount()
}

// Retain the size of the record with the given count
// It will not change the size of the record, but it will change valid column if the count is greater than the size of the record
func (r *Record) Retain(count int) {
	r.tableIndex.Truncate(count)
}
func (r *Record) Merge(other *Record) *Record {
	if other == nil || other.Size() <= 0 {
		return r
	}
	if r.Size() == 0 {
		return other
	}
	r.cacheData = nil
	other.cacheData = nil
	size := r.tableIndex.RowCount()
	otherSize := other.tableIndex.RowCount()

	indexes := make([]uint32, size+otherSize)
	copy(indexes, r.tableIndex.Indexes())
	columnSize := -1
	r.Lock()
	defer r.Unlock()
	columnNamesMap := make(map[string]bool, len(r.columnData))
	for name, column := range r.columnData {
		if column == nil {
			continue
		}

		if columnSize == -1 {
			columnSize = column.Size()
		}
		columnNamesMap[name] = true
		switch recordColumn := column.(type) {
		case *Column[string]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([]string, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[string]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
					//recordColumn.data = append(recordColumn.data, otherRecordColumn.data...)
				} else {
					data := make([]string, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[int]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([]int, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[int]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}

				} else {
					data := make([]int, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[int32]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([]int32, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[int32]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([]int32, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[int64]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([]int64, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[int64]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([]int64, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[float32]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([]float32, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[float32]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([]float32, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[float64]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([]float64, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[float64]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([]float64, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[bool]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([]bool, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[bool]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([]bool, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[[]string]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([][]string, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[[]string]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([][]string, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[[]float32]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([][]float32, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[[]float32]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([][]float32, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[[]int32]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([][]int32, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[[]int32]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([][]int32, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[[]int64]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([][]int64, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[[]int64]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([][]int64, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[[]int]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([][]int, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[[]int]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([][]int, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		case *Column[[]float64]:
			if otherColumn, exist := other.columnData[name]; !exist {
				data := make([][]float64, otherSize)
				recordColumn.data = append(recordColumn.data, data...)
			} else {
				if otherRecordColumn, ok := otherColumn.(*Column[[]float64]); ok {
					for i := 0; i < otherSize; i++ {
						recordColumn.data = append(recordColumn.data, otherRecordColumn.SafeGet(other.tableIndex.SafeGet(i)))
					}
				} else {
					data := make([][]float64, otherSize)
					recordColumn.data = append(recordColumn.data, data...)
				}
			}
		}
	}

	totalSize := columnSize + otherSize
	for name, column := range other.columnData {
		if column == nil {
			continue
		}
		if columnNamesMap[name] {
			continue
		}
		switch recordColumn := column.(type) {
		case *Column[string]:
			newColumn := NewColumn[string](totalSize)

			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[int]:
			newColumn := NewColumn[int](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[int32]:
			newColumn := NewColumn[int32](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[int64]:
			newColumn := NewColumn[int64](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[float32]:
			newColumn := NewColumn[float32](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[float64]:
			newColumn := NewColumn[float64](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[bool]:
			newColumn := NewColumn[bool](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[[]string]:
			newColumn := NewColumn[[]string](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[[]float32]:
			newColumn := NewColumn[[]float32](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[[]int32]:
			newColumn := NewColumn[[]int32](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[[]int64]:
			newColumn := NewColumn[[]int64](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[[]int]:
			newColumn := NewColumn[[]int](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		case *Column[[]float64]:
			newColumn := NewColumn[[]float64](totalSize)
			for i := 0; i < otherSize; i++ {
				//recordColumn.data = append(recordColumn.data)
				newColumn.SetValue(i+columnSize, recordColumn.SafeGet(other.tableIndex.SafeGet(i)))
			}
			r.columnData[name] = newColumn
		}
	}
	for i := size; i < size+otherSize; i++ {
		indexes[i] = uint32(columnSize)
		columnSize++
	}
	r.tableIndex.Rebuild(size+otherSize, indexes)

	return r
}
func (r *Record) Filter(name string) *Record {
	r.RLock()
	column, ok := r.columnData[name]
	r.RUnlock()
	if !ok {
		return r
	}
	size := r.tableIndex.RowCount()
	removed := 0
	uniqueMap := make(map[string]struct{}, size)
	var recordIndex int
	for i := 0; i < size; i++ {
		recordIndex = r.tableIndex.SafeGet(i)
		v, _ := column.Get(recordIndex)
		val := common.ToString(v, "")
		if _, ok := uniqueMap[val]; ok {
			r.tableIndex.MarkRemoved(i)
			removed++
		} else {
			uniqueMap[val] = struct{}{}
		}
	}
	if removed > 0 {
		r.tableIndex.EndRemove()
	}

	return r
}
func (r *Record) FilterByColumnValue(name string, f func(v any) bool) *Record {
	r.RLock()
	column, ok := r.columnData[name]
	r.RUnlock()
	if !ok {
		return r
	}
	size := r.tableIndex.RowCount()
	removed := 0
	var recordIndex int
	for i := 0; i < size; i++ {
		recordIndex = r.tableIndex.SafeGet(i)
		v, _ := column.Get(recordIndex)
		if !f(v) {
			r.tableIndex.MarkRemoved(i)
			removed++
		}
	}
	if removed > 0 {
		r.tableIndex.EndRemove()
	}

	return r
}

// FilterByValues 根据给定的过滤函数对记录进行过滤操作
// 参数:
//   - f: 过滤函数，接收一个map[string]any类型的参数，返回bool值表示是否保留该记录
//
// 返回值:
//   - *Record: 返回当前记录对象的指针，支持链式调用
//
// 该函数会遍历所有记录，对每条记录执行过滤函数，如果返回false则标记为删除
// 最终会更新索引并返回处理后的记录对象
func (r *Record) FilterByValues(f func(m map[string]any) bool) *Record {
	r.RLock()
	defer r.RUnlock()

	size := r.tableIndex.RowCount()
	removed := 0
	var recordIndex int

	// 初始化缓存数据
	if r.cacheData == nil {
		r.cacheData = make(map[int]map[string]any, size)
	}

	var (
		m     map[string]any
		exist bool
	)

	// 遍历所有记录进行过滤处理
	for i := 0; i < size; i++ {
		recordIndex = r.tableIndex.SafeGet(i)
		m, exist = r.cacheData[recordIndex]

		// 如果缓存中不存在该记录，则从列数据中构建记录map
		if !exist {
			m = make(map[string]any, len(r.columnData))
			for name, column := range r.columnData {
				v, _ := column.Get(recordIndex)
				m[name] = v
			}
			r.cacheData[recordIndex] = m
		}

		// 执行过滤函数，如果返回false则标记为删除
		if !f(m) {
			r.tableIndex.MarkRemoved(i)
			removed++
		}
	}

	// 如果有记录被删除，则结束删除操作并更新索引
	if removed > 0 {
		r.tableIndex.EndRemove()
	}

	return r
}
func (r *Record) ColumnValues(columnName string) (ret []any) {
	r.RLock()
	column := r.columnData[columnName]
	r.RUnlock()
	if column == nil {
		return nil
	}

	size := r.tableIndex.RowCount()
	ret = make([]any, 0, size)
	for i := 0; i < size; i++ {
		recordIndex := r.tableIndex.SafeGet(i)
		if v, err := column.Get(recordIndex); err == nil {
			ret = append(ret, v)
		}
	}
	return
}

func (r *Record) ApplyValues(columnName string, f func(m map[string]any) any) *Record {

	columnData, ok := r.columnData[columnName]
	if !ok {
		return r
	}

	size := r.tableIndex.RowCount()
	var recordIndex int

	// 初始化缓存数据
	if r.cacheData == nil {
		r.cacheData = make(map[int]map[string]any, size)
	}

	var (
		m     map[string]any
		exist bool
	)

	// 遍历所有记录进行过滤处理
	for i := 0; i < size; i++ {
		recordIndex = r.tableIndex.SafeGet(i)
		m, exist = r.cacheData[recordIndex]

		// 如果缓存中不存在该记录，则从列数据中构建记录map
		if !exist {
			m = make(map[string]any, len(r.columnData))
			for name, column := range r.columnData {
				v, _ := column.Get(recordIndex)
				m[name] = v
			}
			r.cacheData[recordIndex] = m
		}

		v := f(m)
		switch val := v.(type) {
		case string:
			columnData.(*Column[string]).SetValue(recordIndex, val)
		case int32:
			columnData.(*Column[int32]).SetValue(recordIndex, val)
		case int64:
			columnData.(*Column[int64]).SetValue(recordIndex, val)
		case float32:
			columnData.(*Column[float32]).SetValue(recordIndex, val)
		case float64:
			columnData.(*Column[float64]).SetValue(recordIndex, val)
		case bool:
			columnData.(*Column[bool]).SetValue(recordIndex, val)
		case []string:
			columnData.(*Column[[]string]).SetValue(recordIndex, val)
		case []int32:
			columnData.(*Column[[]int32]).SetValue(recordIndex, val)
		case []int64:
			columnData.(*Column[[]int64]).SetValue(recordIndex, val)
		case []float32:
			columnData.(*Column[[]float32]).SetValue(recordIndex, val)
		case []float64:
			columnData.(*Column[[]float64]).SetValue(recordIndex, val)
		default:
			if str := common.ToString(val, ""); str != "" {
				columnData.(*Column[string]).SetValue(recordIndex, str)
			}
		}
		m[columnName] = v
	}

	return r
}

// ColumnValuesString 根据列名获取该列所有有效记录的字符串表示
// 参数:
//
//	columnName: 要获取数据的列名
//
// 返回值:
//
//	ret: 该列所有有效记录的字符串切片，如果列不存在则返回nil
func (r *Record) ColumnValuesString(columnName string) (ret []string) {
	r.RLock()
	column := r.columnData[columnName]
	r.RUnlock()

	if column == nil {
		return nil
	}

	size := r.tableIndex.RowCount()
	ret = make([]string, 0, size)

	for i := 0; i < size; i++ {
		recordIndex := r.tableIndex.SafeGet(i)
		if v, err := column.Get(recordIndex); err == nil {
			if vv := common.ToString(v, ""); vv != "" {
				ret = append(ret, vv)
			}
		}
	}
	return
}
func (r *Record) ColumnValuesFloat64(columnName string) (ret []float64) {
	r.RLock()
	column := r.columnData[columnName]
	r.RUnlock()

	if column == nil {
		return nil
	}

	size := r.tableIndex.RowCount()
	ret = make([]float64, 0, size)

	for i := 0; i < size; i++ {
		recordIndex := r.tableIndex.SafeGet(i)
		if v, err := column.Get(recordIndex); err == nil {
			ret = append(ret, common.ToFloat(v, 0))
		}
	}
	return
}

func (r *Record) Random() {
	indexes := r.tableIndex.Indexes()
	rand.Shuffle(len(indexes), func(i, j int) {
		indexes[i], indexes[j] = indexes[j], indexes[i]
	})

	r.tableIndex.Rebuild(len(indexes), indexes)
}

func (r *Record) String() string {
	r.RLock()
	defer r.RUnlock()
	buf := bytes.NewBuffer(nil)
	size := r.tableIndex.RowCount()

	buf.WriteString(fmt.Sprintf("count:%d\t", size))
	var fieldNames []string
	for name, column := range r.columnData {
		if column == nil {
			continue
		}
		fieldNames = append(fieldNames, name)
	}
	buf.WriteString(fmt.Sprintf("names:%v\t", fieldNames))
	for _, name := range fieldNames {
		column := r.GetColumn(name)
		var values []any
		for i := 0; i < size; i++ {
			index := r.tableIndex.SafeGet(i)
			v, _ := column.Get(index)
			values = append(values, v)
		}

		buf.WriteString(fmt.Sprintf("%s:%v\t", name, values))
	}

	return buf.String()
}

type sortItem struct {
	index int
	value any
}

type sortItemSlice []*sortItem

func (s sortItemSlice) Less(i, j int) bool {
	switch s[i].value.(type) {
	case float32:
		return s[i].value.(float32) < s[j].value.(float32)
	case float64:
		return s[i].value.(float64) < s[j].value.(float64)
	case int:
		return s[i].value.(int) < s[j].value.(int)
	case int32:
		return s[i].value.(int32) < s[j].value.(int32)
	case int64:
		return s[i].value.(int64) < s[j].value.(int64)
	case string:
		return s[i].value.(string) < s[j].value.(string)
	default:
		return false
	}
}

func (s sortItemSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortItemSlice) Len() int {
	return len(s)
}
