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

var (
	ErrStyleHeader                      = errors.New("header style error")
	ErrStyleBody                        = errors.New("body style error")
	ErrStyleFooter                      = errors.New("footer style error")
	ErrStyleCornerIsUndefined           = errors.New("corder style not defined")
	ErrStyleBorderHorizontalIsUndefined = errors.New("horizontal border style is not defined")
	ErrStyleBorderVerticalIsUndefined   = errors.New("vertical border style is not defined")
)
