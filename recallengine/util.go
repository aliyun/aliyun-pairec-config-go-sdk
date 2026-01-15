package recallengine

import (
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/recallengine/recallenginefb"
	flatbuffers "github.com/google/flatbuffers/go"
)

func UnSerializeRecord(data []byte) *Record {
	matchRecords := recallenginefb.GetRootAsMatchRecords(data, 0)
	record := NewRecord(int(matchRecords.DocCount()))
	size := record.Size()
	for i := 0; i < matchRecords.RecordColumnsLength(); i++ {
		name := string(matchRecords.FieldName(i))
		column := new(recallenginefb.FieldValueColumnTable)
		matchRecords.RecordColumns(column, i)
		switch column.FieldValueColumnType() {
		case recallenginefb.FieldValueColumnStringValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.StringValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[string](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.SetValue(j, string(col.Value(j)))
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnInt32ValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.Int32ValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[int32](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.SetValue(j, col.Value(j))
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnInt64ValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.Int64ValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[int64](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.SetValue(j, col.Value(j))
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnFloatValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.FloatValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[float32](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.SetValue(j, col.Value(j))
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnDoubleValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.DoubleValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[float64](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.SetValue(j, col.Value(j))
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnBoolValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.BoolValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[bool](size)
			for j := 0; j < col.ValueLength(); j++ {
				recordColumn.SetValue(j, col.Value(j))
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnMultiFloatValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.MultiFloatValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[[]float32](size)
			for j := 0; j < col.ValueLength(); j++ {
				multiValue := new(recallenginefb.MultiFloatValue)
				col.Value(multiValue, j)
				list := make([]float32, multiValue.ValueLength())
				for i := 0; i < multiValue.ValueLength(); i++ {
					list[i] = multiValue.Value(i)
				}
				recordColumn.SetValue(j, list)
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnMultiDoubleValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.MultiDoubleValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[[]float64](size)
			for j := 0; j < col.ValueLength(); j++ {
				multiValue := new(recallenginefb.MultiDoubleValue)
				col.Value(multiValue, j)
				list := make([]float64, multiValue.ValueLength())
				for i := 0; i < multiValue.ValueLength(); i++ {
					list[i] = multiValue.Value(i)
				}
				recordColumn.SetValue(j, list)
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnMultiInt32ValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.MultiInt32ValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[[]int32](size)
			for j := 0; j < col.ValueLength(); j++ {
				multiValue := new(recallenginefb.MultiInt32Value)
				col.Value(multiValue, j)
				list := make([]int32, multiValue.ValueLength())
				for i := 0; i < multiValue.ValueLength(); i++ {
					list[i] = multiValue.Value(i)
				}
				recordColumn.SetValue(j, list)
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnMultiInt64ValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.MultiInt64ValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[[]int64](size)
			for j := 0; j < col.ValueLength(); j++ {
				multiValue := new(recallenginefb.MultiInt64Value)
				col.Value(multiValue, j)
				list := make([]int64, multiValue.ValueLength())
				for i := 0; i < multiValue.ValueLength(); i++ {
					list[i] = multiValue.Value(i)
				}
				recordColumn.SetValue(j, list)
			}
			record.SetColumn(name, recordColumn)
		case recallenginefb.FieldValueColumnMultiStringValueColumn:
			unionTab := &flatbuffers.Table{}
			col := recallenginefb.MultiStringValueColumn{}
			column.FieldValueColumn(unionTab)
			col.Init(unionTab.Bytes, unionTab.Pos)
			recordColumn := NewColumn[[]string](size)
			for j := 0; j < col.ValueLength(); j++ {
				multiValue := new(recallenginefb.MultiStringValue)
				col.Value(multiValue, j)
				list := make([]string, multiValue.ValueLength())
				for i := 0; i < multiValue.ValueLength(); i++ {
					list[i] = string(multiValue.Value(i))
				}
				recordColumn.SetValue(j, list)
			}
			record.SetColumn(name, recordColumn)
		}
	}

	return record
}
