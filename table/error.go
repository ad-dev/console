package table

import (
	"errors"
	"fmt"
)

var ErrNoRows = errors.New("table has no rows")

type ErrUnsupportedCell struct {
	Index int
}

func (e *ErrUnsupportedCell) Error() string {
	return fmt.Sprintf("unsupported cell at index %d", e.Index)
}
