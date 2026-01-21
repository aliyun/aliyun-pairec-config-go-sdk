package recallengine

import (
	"errors"
)

type IColumn interface {
	Get(index int) (any, error)
	//SetValue(index int, d any) error
	Size() int
}

var _ IColumn = (*Column[any])(nil)

type Column[T any] struct {
	data []T
}

func NewColumn[T any](size int) *Column[T] {
	column := &Column[T]{
		data: make([]T, size),
	}

	return column
}
func NewColumnWithValue[T any](size int, d T) *Column[T] {
	column := &Column[T]{
		data: make([]T, size),
	}

	for i := 0; i < size; i++ {
		column.data[i] = d
	}
	return column
}
func NewColumnWithArray[T any](arrs []T) *Column[T] {
	column := &Column[T]{
		data: make([]T, len(arrs)),
	}

	copy(column.data, arrs)
	return column
}

func (c *Column[T]) Get(index int) (any, error) {
	if index >= len(c.data) {
		var zero T
		return zero, errors.New("index out of range")
	}
	return c.data[index], nil
}
func (c *Column[T]) SafeGet(index int) T {
	return c.data[index]
}
func (c *Column[T]) SetValue(index int, d T) error {
	if index >= len(c.data) {
		return errors.New("index out of range")
	}
	c.data[index] = d
	return nil

}

func (c *Column[T]) SetData(d []T) {
	size := len(c.data)
	if len(d) < size {
		size = len(d)
	}
	for i := 0; i < size; i++ {
		c.data[i] = d[i]
	}
}
func (c *Column[T]) Size() int {
	return len(c.data)
}

var _ IColumn = (*ConstColumn[any])(nil)

type ConstColumn[T any] struct {
	data T
	size int
}

func NewConstColumn[T any](size int, d T) *ConstColumn[T] {
	column := &ConstColumn[T]{
		data: d,
		size: size,
	}

	return column
}

// Get implements IColumn.
func (c *ConstColumn[T]) Get(index int) (any, error) {
	return c.data, nil
}

// Size implements IColumn.
func (c *ConstColumn[T]) Size() int {
	return c.size
}

func (c *ConstColumn[T]) SafeGet(index int) T {
	return c.data
}
