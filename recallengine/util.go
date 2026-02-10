package recallengine

import (
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/recallengine/recallenginefb"
	flatbuffers "github.com/google/flatbuffers/go"
)

func UnSerializeRecord(data []byte) *Record {
	matchRecords := recallenginefb.GetRootAsMatchRecords(data, 0)
	record := NewRecord(int(matchRecords.DocCount()))
	size := record.Size()
	column := new(recallenginefb.FieldValueColumnTable)
	unionTab := &flatbuffers.Table{}
	for i := 0; i < matchRecords.RecordColumnsLength(); i++ {
		name := string(matchRecords.FieldName(i))
		matchRecords.RecordColumns(column, i)
		switch column.FieldValueColumnType() {
		case recallenginefb.FieldValueColumnStringValueColumn:
			col := recallenginefb.StringValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[string](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.data[j] = string(col.Value(j))
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnInt32ValueColumn:
			col := recallenginefb.Int32ValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[int32](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.data[j] = col.Value(j)
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnInt64ValueColumn:
			col := recallenginefb.Int64ValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[int64](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.data[j] = col.Value(j)
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnFloatValueColumn:
			col := recallenginefb.FloatValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[float32](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.data[j] = col.Value(j)
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnDoubleValueColumn:
			col := recallenginefb.DoubleValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[float64](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.data[j] = col.Value(j)
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnBoolValueColumn:
			col := recallenginefb.BoolValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[bool](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.data[j] = col.Value(j)
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnMultiFloatValueColumn:
			col := recallenginefb.MultiFloatValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[[]float32](size)
			multiValue := new(recallenginefb.MultiFloatValue)
			for j := 0; j < col.ValueLength(); j++ {
				col.Value(multiValue, j)
				length := multiValue.ValueLength()
				list := make([]float32, length)
				for i := 0; i < length; i++ {
					list[i] = multiValue.Value(i)
				}
				recordColumn.data[j] = list
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnMultiDoubleValueColumn:
			col := recallenginefb.MultiDoubleValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[[]float64](size)
			multiValue := new(recallenginefb.MultiDoubleValue)
			for j := 0; j < col.ValueLength(); j++ {
				col.Value(multiValue, j)
				length := multiValue.ValueLength()
				list := make([]float64, length)
				for i := 0; i < length; i++ {
					list[i] = multiValue.Value(i)
				}
				recordColumn.data[j] = list
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnMultiInt32ValueColumn:
			col := recallenginefb.MultiInt32ValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[[]int32](size)
			multiValue := new(recallenginefb.MultiInt32Value)
			for j := 0; j < col.ValueLength(); j++ {
				col.Value(multiValue, j)
				length := multiValue.ValueLength()
				list := make([]int32, length)
				for i := 0; i < length; i++ {
					list[i] = multiValue.Value(i)
				}
				recordColumn.data[j] = list
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnMultiInt64ValueColumn:
			col := recallenginefb.MultiInt64ValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[[]int64](size)
			multiValue := new(recallenginefb.MultiInt64Value)
			for j := 0; j < col.ValueLength(); j++ {
				col.Value(multiValue, j)
				length := multiValue.ValueLength()
				list := make([]int64, length)
				for i := 0; i < length; i++ {
					list[i] = multiValue.Value(i)
				}
				recordColumn.data[j] = list
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnMultiStringValueColumn:
			col := recallenginefb.MultiStringValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[[]string](size)
			multiValue := new(recallenginefb.MultiStringValue)
			for j := 0; j < col.ValueLength(); j++ {
				col.Value(multiValue, j)
				length := multiValue.ValueLength()
				list := make([]string, length)
				for i := 0; i < length; i++ {
					list[i] = string(multiValue.Value(i))
				}
				recordColumn.data[j] = list
			}
			record.SetColumn(name, recordColumn)
		}
	}

	return record
}
