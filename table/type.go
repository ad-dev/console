package table

const (
	Equals byte = iota
	NotEquals
	LowerThan
	GreaterThan
)

type CellCondition struct {
	RowIndex  int
	CellIndex int
	Operator  byte
	Value     string
}

type CellFormatterCallback = func(i int, text string) string
