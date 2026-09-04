package table

import (
	"os"
	"sync"

	"github.com/ad-dev/console"
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

type Hex = uint64

type AsciiTable struct {
	sync.Mutex
	rows                [][]string
	header              []string
	footer              []string
	colWidths           []uint
	dest                *os.File
	cellWidth           uint
	addRowDiv           bool
	defaultPadding      byte
	innerCellSpacing    byte
	paddings            []byte
	truncateCells       []bool
	truncateAllCells    bool
	cellFormatters      formattersMap
	styleHeader         console.Style
	styleBody           console.Style
	styleFooter         console.Style
	currentTheme        console.Theme
	customCellWidths    map[int]map[int]int
	customElements      map[location]console.TextElement
	borderInbetweenRows []console.TableCell
}

type (
	AnyValue    struct{ v any }
	ValueChange struct{ v1, v2 any }
)
