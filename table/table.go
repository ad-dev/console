package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type padding byte

type Hex = uint64

const (
	PAD_LEFT  padding = 1
	PAD_RIGHT padding = 2
)

type AsciiTable struct {
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
	for _, r := range t.rows {
		if len(r) > maxRowLen {
			maxRowLen = len(r)
		}
	}
	return maxRowLen
}
func (t *AsciiTable) getMaxRowCellWidth(row []string) int {
	if len(row) == 0 {
		return 0
	}
	maxWidth := len(row[0])
	for _, c := range row {
		if len(c) > maxWidth {
			maxWidth = len(c)
		}
	}
	return maxWidth
}
func (t *AsciiTable) getMaxCellWidth() int {
	maxWidth := 0
	m := 0
	for _, row := range t.rows {
		m = t.getMaxRowCellWidth(row)
		if m > maxWidth {
			maxWidth = m
		}
	}
	m = t.getMaxRowCellWidth(t.header)
	if m > maxWidth {
		maxWidth = m
	}
	m = t.getMaxRowCellWidth(t.footer)
	if m > maxWidth {
		maxWidth = m
	}

	return maxWidth
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
		return fmt.Sprintf("%s...", strings.TrimSpace(str[:t.cellWidth-1]))
	}
	return str
}

func (t *AsciiTable) displayRow(row []string, cellWidths []uint, p padding) {
	var pd padding
	fmt.Fprint(t.dest, "|")
	for j := range row {
		cellWidth := cellWidths[0]
		if len(cellWidths) == len(row) {
			cellWidth = cellWidths[j]
		}
		pd = p

		if j < len(t.paddings) && (t.paddings[j] == PAD_LEFT || t.paddings[j] == PAD_RIGHT) {
			pd = t.paddings[j]
		}

		if pd == PAD_LEFT {
			fmt.Fprintf(t.dest, "%-"+strconv.Itoa(int(cellWidth))+"s |", t.formatCell(j, row[j]))
		} else {
			fmt.Fprintf(t.dest, "%"+strconv.Itoa(int(cellWidth))+"s |", t.formatCell(j, row[j]))
		}

	}
	fmt.Fprintln(t.dest)
}

func (t *AsciiTable) displayBorder(rowLen int, cellWidths []uint) {
	fmt.Fprint(t.dest, "+")
	for i := 0; i < rowLen; i++ {
		cellWidth := cellWidths[0]
		if len(cellWidths) == rowLen {
			cellWidth = cellWidths[i]
		}
		fmt.Fprintf(t.dest, "%"+strconv.Itoa(int(cellWidth))+"s+", strings.Repeat("-", int(cellWidth+1)))

	}
	fmt.Fprintln(t.dest)
}

func (t *AsciiTable) getColWidths() []uint {

	var cellWidths = make([]uint, t.getMaxRowLen())
	for i := range cellWidths {
		cellWidths[i] = t.cellWidth
		for _, row := range t.rows {
			cw := len(row[i])
			if cw > int(cellWidths[i]) {
				cellWidths[i] = uint(cw)
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

	t.displayBorder(maxRowLen, colWidths)
	if len(t.header) > 0 {
		t.displayRow(t.header, colWidths, t.defaultPadding)
		t.displayBorder(maxRowLen, colWidths)
	}
	for i := range t.rows {
		t.displayRow(
			t.alignRow(t.rows[i], maxRowLen, PAD_RIGHT),
			colWidths, t.defaultPadding)
		if t.addRowDiv && (i < len(t.rows)-1) {
			t.displayBorder(maxRowLen, colWidths)
		}
	}
	if len(t.footer) > 0 {
		t.displayBorder(maxRowLen, colWidths)
		t.displayRow(t.footer, colWidths, t.defaultPadding)
	}
	t.displayBorder(maxRowLen, colWidths)
	return nil
}

func (t *AsciiTable) SetDefaultPadding(p padding) {
	t.defaultPadding = p

}

func (t *AsciiTable) SetTruncateAllCells(flag bool) {
	t.truncateAllCells = flag
	print("yo", t.truncateAllCells)

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

func New(cellWidth uint, addRowDiv bool, dest *os.File) *AsciiTable {
	return &AsciiTable{cellWidth: cellWidth, addRowDiv: addRowDiv, dest: dest, defaultPadding: PAD_RIGHT}
}
