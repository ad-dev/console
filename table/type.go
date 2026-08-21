package table

import "github.com/ad-dev/console"

const (
	Equals byte = iota
	NotEquals
	LowerThan
	GreaterThan
)

type CustomCellWidth struct {
	Content string
	Width   int
}

type location struct {
	row int
	col int
}

type formattersMap = map[location]map[*console.TableCellFormatterCallback][]console.TableCellCondition
