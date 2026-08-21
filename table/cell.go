package table

import "github.com/ad-dev/console"

type cell[T any] struct {
	row   int
	col   int
	value T
}

func (c *cell[T]) SetRow(i int) console.TableCell {
	c.row = i
	return c
}

func (c *cell[T]) SetColumn(i int) console.TableCell {
	c.col = i
	return c
}

func (c *cell[T]) SetValue(v any) console.TableCell {
	c.value = v.(T)
	return c
}

func (c *cell[T]) Row() int {
	return c.row
}

func (c *cell[T]) Column() int {
	return c.col
}
func (c *cell[T]) Value() any {
	return c.value
}

func NewCell[T any](r, c int, v T) console.TableCell {
	return &cell[T]{row: r, col: c, value: v}
}
