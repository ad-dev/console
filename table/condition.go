package table

import "github.com/ad-dev/console"

type condition[T any] struct {
	*cell[T]
	operator byte
}

func (c *condition[T]) SetOperator(o byte) console.TableCellCondition {
	c.operator = o
	return c
}

func (c *condition[T]) Operator() byte {
	return c.operator
}

func (c *condition[T]) Set(row, col int, o byte, v any) console.TableCellCondition {
	c.operator = o

	c.cell.col = col
	c.cell.row = row
	c.cell.value = v.(T)

	return c
}

func NewCellCondition[T any](r, c int, o byte, v T) console.TableCellCondition {
	return &condition[T]{cell: &cell[T]{row: r, col: c, value: v}, operator: o}
}
