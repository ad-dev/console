package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type padding byte

type Style = map[uint]string

type Hex = uint64

const (
	PAD_LEFT  padding = 2
	PAD_RIGHT padding = 4

	ALIGN_TOP    padding = 8
	ALIGN_BOTTOM padding = 16

	DEFAULT_STYLE_CORNER            = "+"
	DEFAULT_STYLE_BORDER_HORIZONTAL = "-"
	DEFAULT_STYLE_BORDER_VERTICAL   = "|"

	STYLE_CORNER            = 0
	STYLE_BORDER_HORIZONTAL = 1
	STYLE_BORDER_VERTICAL   = 2
)

type AsciiTable struct {
	sync.Mutex
	rows             [][]string
	header           []string
	footer           []string
	colWidths        []uint
	dest             *os.File
	cellWidth        uint
	addRowDiv        bool
	defaultPadding   padding
	paddings         []padding
	truncateCells    []bool
	truncateAllCells bool
	cellFormatters   []map[*CellFormatterCallback][]CellCondition
	styleHeader      Style
	styleBody        Style
	styleFooter      Style
	customCellWidths map[int]map[int]int
}

func (t *AsciiTable) ClearRows() {
	t.rows = make([][]string, 0)
}

func (t *AsciiTable) AddRow(row []string) {
	t.rows = append(t.rows, row)
}

func (t *AsciiTable) convertAnyRow(row []any) ([]string, error) {
	var err error
	r := make([]string, len(row))
	cri := len(t.rows) - 1
	for i, c := range row {
		switch c := c.(type) {
		case string:
			r[i] = c
		case Hex:
			r[i] = fmt.Sprintf("0x%X", c)
		case byte:
			r[i] = fmt.Sprintf("0x%02X", c)
		case []byte:
			if len(c) > 0 {
				r[i] = fmt.Sprintf("% 02X", c)
			} else {
				r[i] = ""
			}
		case int:
			r[i] = strconv.Itoa(int(c))
		case int64:
			r[i] = strconv.Itoa(int(c))
		case int32:
			r[i] = strconv.Itoa(int(c))
		case uint32:
			r[i] = strconv.Itoa(int(c))
		case int16:
			r[i] = strconv.Itoa(int(c))
		case uint:
			r[i] = strconv.Itoa(int(c))
		case uint16:
			r[i] = strconv.Itoa(int(c))
		case uintptr:
			r[i] = strconv.Itoa(int(c))
		case CustomCellWidth:
			r[i] = c.Content
			t.Lock()

			if t.customCellWidths == nil {
				t.customCellWidths = make(map[int]map[int]int)
			}
			if t.customCellWidths[cri] == nil {

				t.customCellWidths[cri] = make(map[int]int)
			}
			t.customCellWidths[cri][i] = c.Width
			t.Unlock()

		default:
			err = errors.Join(err, &ErrUnsupportedCell{Index: i})
		}
	}
	return r, err
}

func (t *AsciiTable) AddAnyRow(row []any) error {
	r, err := t.convertAnyRow(row)
	if err != nil {
		return err
	}
	t.rows = append(t.rows, r)
	return nil
}

func (t *AsciiTable) AddHeader(row []string) {
	t.header = row
}

func (t *AsciiTable) AddFooter(row []string) {
	t.footer = row
}

func (t *AsciiTable) AddAnyFooter(row []any) error {
	r, err := t.convertAnyRow(row)
	if err != nil {
		return err
	}
	t.footer = r
	return nil
}

func (t *AsciiTable) AddAnyHeader(row []any) error {
	r, err := t.convertAnyRow(row)
	if err != nil {
		return err
	}
	t.header = r
	return nil
}

func (t *AsciiTable) alignRow(row []string, maxLen int, p padding) []string {
	if maxLen > len(row) {
		padding := maxLen - len(row)
		emptyCells := make([]string, padding)
		switch p {
		case PAD_RIGHT:
			row = append(row, emptyCells...)
		case PAD_LEFT:
			row = append(emptyCells, row...)
		}
	}
	return row
}

func (t *AsciiTable) getMaxRowLen() int {
	maxRowLen := 0
	l := 0
	for _, r := range t.rows {
		l = len(r)
		if l > maxRowLen {
			maxRowLen = l
		}
	}
	return maxRowLen
}

func (t *AsciiTable) formatCell(j int, str string) string {

	if t.cellWidth < 1 || len(str) < int(t.cellWidth)-1 {
		return str
	}

	truncate := false

	if t.truncateAllCells {
		truncate = true
	}

	if j < len(t.truncateCells) {
		truncate = t.truncateCells[j]
	}

	if truncate {
		str = fmt.Sprintf("%s...", strings.TrimSpace(str[:t.cellWidth-1]))
	}
	if j < len(t.cellFormatters) {
		if formatters := t.cellFormatters[j]; formatters != nil {
			for formatter, conditions := range formatters {
				f := *formatter
				if t.conditionsAreMet(conditions...) {
					str = f(j, str)
				}

			}
		}

	}
	return str
}

func (t *AsciiTable) displayRow(row []string, cellWidths []uint, p padding, s Style) {
	var pd padding
	noMultiCellsInARow := !t.doesRowContainMultilineCells(row)
	if noMultiCellsInARow {
		fmt.Fprint(t.dest, s[STYLE_BORDER_VERTICAL])
	}

	linesMax := t.getBiggestMultilineCell(row)
	if linesMax > 1 {
		mlRow := make([]string, t.getMaxRowLen())
		copy(mlRow, row)
		for k := range linesMax {
			for i := range row {
				irs := strings.Split(row[i], "\n")
				if len(irs) > k {
					mlRow[i] = irs[k]
				} else {
					mlRow[i] = ""
				}
			}

			t.displayRow(mlRow, cellWidths, pd, s)
		}
	} else if noMultiCellsInARow {
		for j := range row {
			cellWidth := cellWidths[0]
			if len(cellWidths) == len(row) {
				cellWidth = cellWidths[j]
			}
			pd = p

			if j < len(t.paddings) && (t.paddings[j]&PAD_LEFT == PAD_LEFT || t.paddings[j]&PAD_RIGHT == PAD_RIGHT) {
				p = t.paddings[j]
			}

			if pd&PAD_LEFT == PAD_LEFT {
				fmt.Fprintf(t.dest, "%-"+strconv.Itoa(int(cellWidth))+"s %s", t.formatCell(j, row[j]), s[STYLE_BORDER_VERTICAL])
			} else {
				fmt.Fprintf(t.dest, "%"+strconv.Itoa(int(cellWidth))+"s %s", t.formatCell(j, row[j]), s[STYLE_BORDER_VERTICAL])
			}
		}
	}

	if noMultiCellsInARow {
		fmt.Fprintln(t.dest)
	}
}

func (t *AsciiTable) doesRowContainMultilineCells(row []string) bool {
	for _, cell := range row {
		if strings.Contains(cell, "\n") {
			return true
		}
	}
	return false

}

func (t *AsciiTable) getBiggestMultilineCell(row []string) int {
	var lines []string
	linesMax := 0
	linesCount := 0

	for _, cell := range row {
		lines = strings.Split(cell, "\n")
		linesCount = len(lines)
		if linesCount > linesMax {
			linesMax = linesCount
		}
	}
	return linesMax

}

func (t *AsciiTable) displayBorder(rowLen int, cellWidths []uint, style Style) {
	fmt.Fprint(t.dest, style[STYLE_CORNER])
	for i := 0; i < rowLen; i++ {
		cellWidth := cellWidths[0]
		if len(cellWidths) == rowLen {
			cellWidth = cellWidths[i]
		}
		fmt.Fprintf(
			t.dest,
			"%"+strconv.Itoa(int(cellWidth))+"s%s",
			strings.Repeat(style[STYLE_BORDER_HORIZONTAL], int(cellWidth+1)),
			style[STYLE_CORNER],
		)

	}
	fmt.Fprintln(t.dest)
}

func (t *AsciiTable) getColWidths() []uint {

	var cellWidths = make([]uint, t.getMaxRowLen())
	var cw int
	for i := range cellWidths {
		cellWidths[i] = t.cellWidth
		for ri, row := range t.rows {
			if i < len(row) {
				cw = len(row[i])
				t.Lock()
				if ccw, defined := t.customCellWidths[ri][i]; defined {
					cw = ccw
				}
				t.Unlock()
				if cw > int(cellWidths[i]) {
					cellWidths[i] = uint(cw)
				}
			}
		}
		if len(t.header) > i {
			cw := len(t.header[i])
			if cw > int(cellWidths[i]) {
				cellWidths[i] = uint(cw)
			}
		}
		if len(t.footer) > i {
			cw := len(t.footer[i])
			if cw > int(cellWidths[i]) {
				cellWidths[i] = uint(cw)
			}
		}

	}
	return cellWidths
}

func (t *AsciiTable) Display() error {
	if len(t.rows) == 0 {
		return ErrNoRows
	}
	maxRowLen := t.getMaxRowLen()

	if len(t.header) > 0 {
		t.header = t.alignRow(t.header, t.getMaxRowLen(), PAD_RIGHT)
	}

	if len(t.footer) > 0 {
		t.footer = t.alignRow(t.footer, t.getMaxRowLen(), PAD_LEFT)
	}

	colWidths := t.getColWidths()

	t.displayBorder(maxRowLen, colWidths, t.styleHeader)
	if len(t.header) > 0 {
		t.displayRow(t.header, colWidths, t.defaultPadding, t.styleHeader)
		t.displayBorder(maxRowLen, colWidths, t.styleHeader)
	}
	for i := range t.rows {
		t.displayRow(
			t.alignRow(t.rows[i], maxRowLen, PAD_RIGHT),
			colWidths, t.defaultPadding, t.styleBody)
		if t.addRowDiv && (i < len(t.rows)-1) {
			t.displayBorder(maxRowLen, colWidths, t.styleBody)
		}
	}
	if len(t.footer) > 0 {
		t.displayBorder(maxRowLen, colWidths, t.styleFooter)
		t.displayRow(t.footer, colWidths, t.defaultPadding, t.styleFooter)
	}
	t.displayBorder(maxRowLen, colWidths, t.styleFooter)
	return nil
}

func (t *AsciiTable) SetDefaultPadding(p padding) {
	t.defaultPadding = p

}

func (t *AsciiTable) SetTruncateAllCells(flag bool) {
	t.truncateAllCells = flag

}

func (t *AsciiTable) SetCellTruncate(col uint, flag bool) {
	if col >= uint(len(t.truncateCells)) {
		t.truncateCells = make([]bool, col+1)
	}
	t.truncateCells[col] = flag
}

func (t *AsciiTable) SetColumnPadding(col uint, p padding) {
	if col >= uint(len(t.paddings)) {
		t.paddings = make([]padding, col+1)
	}
	t.paddings[col] = p
}

func (t *AsciiTable) getRawOutput() [][]string {
	var output [][]string
	if len(t.rows) == 0 {
		return output
	}
	maxRowLen := t.getMaxRowLen()

	if len(t.header) > 0 {
		t.header = t.alignRow(t.header, t.getMaxRowLen(), PAD_RIGHT)
		output = append(output, t.header)
	}

	for i := range t.rows {
		output = append(output, t.alignRow(t.rows[i], maxRowLen, PAD_RIGHT))
	}

	if len(t.footer) > 0 {
		t.footer = t.alignRow(t.footer, t.getMaxRowLen(), PAD_LEFT)
		output = append(output, t.footer)
	}
	return output
}

func (t *AsciiTable) GetJSON() (string, error) {
	var jsonBytes []byte
	var err error
	jsonBytes, err = json.Marshal(t.getRawOutput())
	return (string)(jsonBytes), err
}

func (t *AsciiTable) GetFormattedJSON() (string, error) {
	var jsonBytes []byte
	var err error
	jsonBytes, err = json.MarshalIndent(t.getRawOutput(), "", "  ")
	return (string)(jsonBytes), err
}

func (t *AsciiTable) PrintFormattedJSON(key string) error {
	jfs, err := t.GetFormattedJSON()
	if err != nil {
		return err
	}
	key = strings.Trim(key, " ")
	fmt.Fprint(t.dest, "```json")
	if key != "" {
		fmt.Fprint(t.dest, " ", key)
	}
	fmt.Fprint(t.dest, "\n", jfs)
	fmt.Fprint(t.dest, "\n```\n")
	return nil

}

func (t *AsciiTable) SetCellFormmatter(index int, cb CellFormatterCallback, conditions ...CellCondition) {
	if index >= len(t.cellFormatters) {
		t.cellFormatters = make([]map[*CellFormatterCallback][]CellCondition, index+1)
	}
	if t.cellFormatters[index] == nil {
		t.cellFormatters[index] = make(map[*CellFormatterCallback][]CellCondition)
	}
	t.cellFormatters[index][&cb] = conditions
}

func (t *AsciiTable) conditionsAreMet(list ...CellCondition) bool {
	if list == nil {
		return true
	}
	flags := make([]bool, len(list))
	for i, c := range list {
		if c.RowIndex >= len(t.rows) {
			continue
		}
		if c.CellIndex >= len(t.rows[c.RowIndex]) {
			continue
		}

		switch c.Operator {
		case Equals:
			flags[i] = c.Value == t.rows[c.RowIndex][c.CellIndex]
		case NotEquals:
			flags[i] = c.Value != t.rows[c.RowIndex][c.CellIndex]
		case GreaterThan:
			v0, e0 := strconv.Atoi(c.Value)
			v1, e1 := strconv.Atoi(t.rows[c.RowIndex][c.CellIndex])
			if e0 == nil && e1 == nil {
				flags[i] = v0 < v1
			}
		case LowerThan:
			v0, e0 := strconv.Atoi(c.Value)
			v1, e1 := strconv.Atoi(t.rows[c.RowIndex][c.CellIndex])
			if e0 == nil && e1 == nil {
				flags[i] = v0 > v1
			}
		}

	}

	for _, f := range flags {
		if !f {
			return false
		}
	}
	return true
}

func (t *AsciiTable) SetHeaderStyle(s Style) error {
	checkStyle(s, ErrStyleHeader)
	t.styleHeader = s

	return nil
}

func (t *AsciiTable) SetBodyStyle(s Style) error {
	checkStyle(s, ErrStyleBody)
	t.styleBody = s

	return nil
}

func (t *AsciiTable) SetFooterStyle(s Style) error {
	checkStyle(s, ErrStyleFooter)
	t.styleFooter = s

	return nil
}

func checkStyle(s Style, err error) error {
	if _, found := s[STYLE_CORNER]; !found {
		return errors.Join(err, ErrStyleCornerIsUndefined)
	}

	if _, found := s[STYLE_BORDER_HORIZONTAL]; !found {
		return errors.Join(err, ErrStyleBorderHorizontalIsUndefined)
	}

	if _, found := s[STYLE_BORDER_VERTICAL]; !found {
		return errors.Join(err, ErrStyleBorderVerticalIsUndefined)
	}
	return nil
}

func New(cellWidth uint, addRowDiv bool, dest *os.File) *AsciiTable {
	return &AsciiTable{
		cellWidth:      cellWidth,
		addRowDiv:      addRowDiv,
		dest:           dest,
		defaultPadding: PAD_RIGHT,
		styleHeader: Style{
			STYLE_CORNER:            DEFAULT_STYLE_CORNER,
			STYLE_BORDER_HORIZONTAL: DEFAULT_STYLE_BORDER_HORIZONTAL,
			STYLE_BORDER_VERTICAL:   DEFAULT_STYLE_BORDER_VERTICAL,
		},
		styleBody: Style{
			STYLE_CORNER:            DEFAULT_STYLE_CORNER,
			STYLE_BORDER_HORIZONTAL: DEFAULT_STYLE_BORDER_HORIZONTAL,
			STYLE_BORDER_VERTICAL:   DEFAULT_STYLE_BORDER_VERTICAL,
		},
		styleFooter: Style{
			STYLE_CORNER:            DEFAULT_STYLE_CORNER,
			STYLE_BORDER_HORIZONTAL: DEFAULT_STYLE_BORDER_HORIZONTAL,
			STYLE_BORDER_VERTICAL:   DEFAULT_STYLE_BORDER_VERTICAL,
		},
	}
}
