package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/ad-dev/console"
	"github.com/ad-dev/console/table/aggregator"
)

func (t *AsciiTable) ClearRows() {
	t.rows = make([][]string, 0)
}

func (t *AsciiTable) AddRow(row []string) {
	t.rows = append(t.rows, row)
}

func (t *AsciiTable) makeCustomCellWidth(row, col int, w int) {
	t.Lock()
	if t.customCellWidths == nil {
		t.customCellWidths = make(map[int]map[int]int)
	}
	if t.customCellWidths[row] == nil {

		t.customCellWidths[row] = make(map[int]int)
	}
	t.customCellWidths[row][col] = w
	t.Unlock()
}

func (t *AsciiTable) addCustomElement(row, col int, elm console.TextElement) {
	t.Lock()
	if t.customElements == nil {
		t.customElements = make(map[location]console.TextElement)
	}
	t.customElements[location{row: row, col: col}] = elm
	t.Unlock()
}

func (t *AsciiTable) getCustomElement(row, col int) console.TextElement {
	t.Lock()
	defer t.Unlock()

	if ce, ok := t.customElements[location{row: row, col: col}]; ok {
		return ce
	}
	return nil
}

func (t *AsciiTable) convertAnyRow(row []any) ([]string, error) {
	var err error
	r := make([]string, len(row))
	cri := len(t.rows)
	for i, c := range row {

		if s, ok := c.(console.TextElement); ok {
			r[i] = s.String()
			t.addCustomElement(cri, i, s)
			continue
		}

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
		case float32:
			r[i] = fmt.Sprintf("%.8f", c)
		case float64:
			r[i] = fmt.Sprintf("%.8f", c)
		case CustomCellWidth:
			r[i] = c.Content
			t.makeCustomCellWidth(cri, i, c.Width)
		case aggregator.Aggregator:
			stop := false
			switch c.Type() {
			case aggregator.Vertical:
				data := make([]string, len(t.rows))
				for ri, row := range t.rows {
					data[ri] = row[i]
				}
				c.SetData(data)
			case aggregator.Horizontal:
				if i == 0 {
					r[i] = "????"
					stop = true
				}
				data := make([]string, len(row)-1)
				for ci := 0; ci < i; ci++ {
					data[ci] = r[ci]
				}

				c.SetData(data)

			default:
				r[i] = aggregator.ErrUnknownResult
				stop = true
				break
			}
			if stop {
				break
			}

			e := c.Calc()
			if e != nil {
				r[i] = e.Error()
				break
			}
			r[i] = c.String()
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

func (t *AsciiTable) ChangeRow(index int, row []any) error {
	r, err := t.convertAnyRow(row)
	if err != nil {
		return err
	}
	if index < len(t.rows) {
		t.rows[index] = r
	}

	return nil
}

func (t *AsciiTable) ChangeHeader(row []any) error {
	r, err := t.convertAnyRow(row)
	if err != nil {
		return err
	}
	t.header = r

	return nil
}

func (t *AsciiTable) ChangeFooter(row []any) error {
	r, err := t.convertAnyRow(row)
	if err != nil {
		return err
	}
	t.footer = r

	return nil
}

func (t *AsciiTable) alignRow(row []string, maxLen int, p byte) []string {
	if maxLen > len(row) {
		padding := maxLen - len(row)
		emptyCells := make([]string, padding)
		switch p {
		case PadRight:
			row = append(row, emptyCells...)
		case PadLeft:
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

		if len(t.header) > maxRowLen {
			maxRowLen = len(t.header)
		}

		if len(t.footer) > maxRowLen {
			maxRowLen = len(t.footer)
		}
	}
	return maxRowLen
}

func (t *AsciiTable) formatCell(l location, str string) string {

	if t.cellWidth < 1 {
		return str
	}

	truncate := false

	if t.truncateAllCells {
		truncate = true
	}

	if l.col < len(t.truncateCells) {
		truncate = t.truncateCells[l.col]
	}

	cellLen := utf8.RuneCountInString(str)
	ce := t.getCustomElement(l.row, l.col)

	if ce == nil {
		if truncate && cellLen > int(t.cellWidth) {
			tr := []rune(str)[:int(t.cellWidth)-utf8.RuneCountInString(truncatedTextIndicator)]
			str = fmt.Sprintf("%s%s", string(tr), truncatedTextIndicator)
		}
	} else {

		if truncate && ce.Len() > int(t.cellWidth) {
			ce.SetTruncateIndicator(truncatedTextIndicator)
			str = fmt.Sprintf(
				"%s%s",
				string(ce.Truncate(int(t.cellWidth))),
				truncatedTextIndicator,
			)
		}
	}

	if lf, found := t.cellFormatters[l]; found {
		for formatter, cb := range lf {
			f := *formatter
			if t.conditionsAreMet(cb...) {
				str = f(NewCell(l.row, l.col, str))
			}
		}
	}

	return str
}

func (t *AsciiTable) displayRow(originalRowIndex int, row []string, cellWidths []uint, p byte, s console.Style) {
	var pd byte
	noMultiCellsInARow := !t.doesRowContainMultilineCells(row)
	if noMultiCellsInARow {
		fmt.Fprint(t.dest, s[StyleBorderVertical])
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

			t.displayRow(originalRowIndex, mlRow, cellWidths, p, s)
		}
	} else if noMultiCellsInARow {
		for j := range row {
			cellWidth := cellWidths[0]
			if len(cellWidths) == len(row) {
				cellWidth = cellWidths[j]
			}
			pd = p

			if j < len(t.paddings) && (t.paddings[j]&PadLeft == PadLeft || t.paddings[j]&PadRight == PadRight) {
				pd = t.paddings[j]
			}

			cellLen := utf8.RuneCountInString(row[j])

			if cl, found := t.customCellWidths[originalRowIndex][j]; found && cl < cellLen {
				cellLen = cl
			}

			if ce := t.getCustomElement(originalRowIndex, j); ce != nil {
				cellLen = ce.Len()

			}

			padding := int(cellWidth) - cellLen + int(t.innerCellSpacing)
			if padding < 0 {
				padding = 0
			}

			spacing := strings.Repeat(defaultPaddingChar, int(t.innerCellSpacing))
			if pd&PadLeft == PadLeft {
				fmt.Fprintf(t.dest, "%s%s%s%s", strings.Repeat(defaultPaddingChar, padding), row[j], spacing, s[StyleBorderVertical])
			} else {
				fmt.Fprintf(t.dest, "%s%s%s%s", spacing, row[j], strings.Repeat(defaultPaddingChar, padding), s[StyleBorderVertical])
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

func (t *AsciiTable) displayBorder(originalRowNo, rowLen int, cellWidths []uint, style console.Style) {
	rc := len(t.rows)

	cornerStyle := style[StyleBorderJointLeft]
	if originalRowNo == 0 {
		cornerStyle = style[StyleCorner]
	} else if originalRowNo == rc {
		cornerStyle = style[StyleCornerBottom]
	}

	cornerStyleRight := style[StyleBorderJoint]
	if originalRowNo == rc {
		cornerStyleRight = style[StyleBorderJointBottom]
	}

	fmt.Fprint(t.dest, cornerStyle)

	for i := 0; i < rowLen; i++ {
		cellWidth := cellWidths[0]
		if len(cellWidths) == rowLen {
			cellWidth = cellWidths[i]
		}

		if i == rowLen-1 && originalRowNo < rc {
			cornerStyleRight = style[StyleBorderJointRight]
		} else if i == rowLen-1 && originalRowNo == rc {
			cornerStyleRight = style[StyleCornerBottomRight]
		}

		if originalRowNo == 0 && i == rowLen-1 {
			cornerStyleRight = style[StyleCornerRight]
		} else if originalRowNo == 0 {
			cornerStyleRight = style[StyleBorderJointTop]
		}

		fmt.Fprintf(
			t.dest,
			"%"+strconv.Itoa(int(cellWidth)+int(t.innerCellSpacing)*2)+"s%s",
			strings.Repeat(style[StyleBorderHorizontal], int(cellWidth)+int(t.innerCellSpacing)*2),
			cornerStyleRight,
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
				if ccw, defined := t.customCellWidths[ri][i]; defined && ccw > cw {
					cw = ccw
				}
				t.Unlock()

				if ce := t.getCustomElement(ri, i); ce != nil {
					cw = ce.Len()
				}

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
		t.header = t.alignRow(t.header, t.getMaxRowLen(), PadRight)
	}

	if len(t.footer) > 0 {
		t.footer = t.alignRow(t.footer, t.getMaxRowLen(), PadLeft)
	}

	for j, col := range t.header {
		t.header[j] = t.formatCell(location{row: HeaderIndex, col: j}, col)
	}

	for i, row := range t.rows {
		for j, col := range row {
			t.rows[i][j] = t.formatCell(location{row: i, col: j}, col)
		}
	}

	for j, col := range t.footer {
		t.footer[j] = t.formatCell(location{row: FooterIndex, col: j}, col)
	}

	colWidths := t.getColWidths()

	t.displayBorder(0, maxRowLen, colWidths, t.styleHeader)
	if len(t.header) > 0 {
		t.displayRow(HeaderIndex, t.header, colWidths, t.defaultPadding, t.styleHeader)
		t.displayBorder(1, maxRowLen, colWidths, t.styleHeader)
	}
	for i := range t.rows {
		t.displayRow(i,
			t.alignRow(t.rows[i], maxRowLen, PadRight),
			colWidths, t.defaultPadding, t.styleBody)
		if (t.addRowDiv || t.addBorderInbetweenRows(i)) && (i < len(t.rows)-1) {
			t.displayBorder(i+1, maxRowLen, colWidths, t.styleBody)
		}
	}
	rc := len(t.rows)
	if len(t.footer) > 0 {
		t.displayBorder(rc-1, maxRowLen, colWidths, t.styleFooter)
		t.displayRow(FooterIndex, t.footer, colWidths, t.defaultPadding, t.styleFooter)
	}
	t.displayBorder(rc, maxRowLen, colWidths, t.styleFooter)
	return nil
}

func (t *AsciiTable) SetDefaultPadding(p byte) {
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

func (t *AsciiTable) SetColumnPadding(col uint, p byte) {
	if col >= uint(len(t.paddings)) {
		t.paddings = make([]byte, col+1)
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
		t.header = t.alignRow(t.header, t.getMaxRowLen(), PadRight)
		output = append(output, t.header)
	}

	for i := range t.rows {
		output = append(output, t.alignRow(t.rows[i], maxRowLen, PadRight))
	}

	if len(t.footer) > 0 {
		t.footer = t.alignRow(t.footer, t.getMaxRowLen(), PadLeft)
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

func (t *AsciiTable) SetCellFormatter(row int, col int, cb console.TableCellFormatterCallback, conditions ...console.TableCellCondition) {
	if t.cellFormatters == nil {
		t.cellFormatters = make(formattersMap)
	}
	l := location{row: row, col: col}
	if t.cellFormatters[l] == nil {
		t.cellFormatters[l] = make(map[*console.TableCellFormatterCallback][]console.TableCellCondition)
	}
	t.cellFormatters[l][&cb] = conditions
}

func (t *AsciiTable) AddCellFormatter(row int, col int, cb console.TableCellFormatterCallback, conditions ...console.TableCellCondition) {
	l := location{row: row, col: col}
	cf := make(formattersMap)
	if t.cellFormatters[l] == nil {
		t.cellFormatters[l] = make(map[*console.TableCellFormatterCallback][]console.TableCellCondition)
	}
	cf[l][&cb] = append(cf[l][&cb], conditions...)
}

func (t *AsciiTable) conditionsAreMet(list ...console.TableCellCondition) bool {
	if list == nil {
		return true
	}
	flags := make([]bool, len(list))
	for i, c := range list {
		if c.Row() >= len(t.rows) {
			continue
		}
		if c.Column() >= len(t.rows[c.Row()]) {
			continue
		}

		switch c.Operator() {
		case Equals:
			flags[i] = c.Value() == t.rows[c.Row()][c.Column()]
		case NotEquals:
			flags[i] = c.Value() != t.rows[c.Row()][c.Column()]
		case GreaterThan:
			v0, e0 := strconv.Atoi(c.Value().(string))
			v1, e1 := strconv.Atoi(t.rows[c.Row()][c.Column()])
			if e0 == nil && e1 == nil {
				flags[i] = v0 < v1
			}
		case LowerThan:
			v0, e0 := strconv.Atoi(c.Value().(string))
			v1, e1 := strconv.Atoi(t.rows[c.Row()][c.Column()])
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

func (t *AsciiTable) SetHeaderStyle(s console.Style) error {
	checkStyle(s, ErrStyleHeader)
	t.styleHeader = s

	return nil
}

func (t *AsciiTable) SetBodyStyle(s console.Style) error {
	checkStyle(s, ErrStyleBody)
	t.styleBody = s

	return nil
}

func (t *AsciiTable) SetFooterStyle(s console.Style) error {
	checkStyle(s, ErrStyleFooter)
	t.styleFooter = s

	return nil
}

func (t *AsciiTable) SetTheme(th console.Theme) error {
	if thm, found := themes[th]; found {
		if x, ok := thm[Header]; ok {
			t.SetHeaderStyle(x)
		} else {
			return ErrStyleHeader
		}

		if x, ok := thm[Body]; ok {
			t.SetBodyStyle(x)
		} else {
			return ErrStyleHeader
		}

		if x, ok := thm[Footer]; ok {
			t.SetFooterStyle(x)
		} else {
			return ErrStyleHeader
		}
		t.currentTheme = th
		return nil
	}
	return ErrThemeNotFound
}

func (t *AsciiTable) SetThemeByName(name string) error {
	for id, n := range themeNames {
		if n == name {
			return t.SetTheme(id)
		}
	}
	return errors.Join(ErrThemeNotFound, errors.New(name))
}

func (t *AsciiTable) Theme() console.Theme {
	return t.currentTheme
}

func (t *AsciiTable) SetOutputDestination(d *os.File) {
	t.dest = d
}

func (t *AsciiTable) GetData() ([]string, [][]string, []string, error) {
	if len(t.rows) == 0 {
		return nil, nil, nil, ErrNoRows
	}

	return t.header, t.rows, t.footer, nil
}

func (t *AsciiTable) AddBorderInbetweenRowsIf(c ...console.TableCell) {
	if c != nil && len(c) > 0 {
		t.borderInbetweenRows = c
	}
}

func (t *AsciiTable) addBorderInbetweenRows(rowIndex int) bool {

	if len(t.borderInbetweenRows) == 0 {
		return false
	}

	fc := false
	afc := true
	var cv0, cv1 string
	var cvf0, cvf1 bool

	for _, tc := range t.borderInbetweenRows {
		fc = false
		switch vt := tc.Value().(type) {
		case ValueChange:
			if tc.Row() != AnyIndex && tc.Column() != AnyIndex {
				if tc.Row() == rowIndex && len(t.rows) > rowIndex && len(t.rows[rowIndex]) > tc.Column() {
					cv0 = t.rows[rowIndex][tc.Column()]
					cvf0 = true
				}

				if tc.Row() == rowIndex && len(t.rows) > rowIndex+1 && len(t.rows[rowIndex+1]) > tc.Column() {
					cv1 = t.rows[rowIndex+1][tc.Column()]
					cvf1 = true
				}

				if vt.v1 == nil && vt.v2 == nil && cvf0 && cvf1 && cv0 != cv1 {
					fc = true
				}

			} else if tc.Column() == AnyIndex {
				if (rowIndex == tc.Row() || tc.Row() == AnyIndex) && len(t.rows) > rowIndex+1 {
					for ci, col := range t.rows[rowIndex] {
						if col != t.rows[rowIndex+1][ci] {
							fc = true
							break
						}
					}
				}
			} else if tc.Row() == AnyIndex && tc.Column() != AnyIndex {
				if len(t.rows) > rowIndex+1 && len(t.rows[rowIndex]) > tc.Column() && len(t.rows[rowIndex+1]) > tc.Column() &&
					t.rows[rowIndex][tc.Column()] != t.rows[rowIndex+1][tc.Column()] {
					fc = true
				}
			}
		default:
		}
		afc = afc && fc
	}

	return afc
}

func checkStyle(s console.Style, err error) error {
	if _, found := s[StyleCorner]; !found {
		return errors.Join(err, ErrStyleCornerIsUndefined)
	}

	if _, found := s[StyleBorderHorizontal]; !found {
		return errors.Join(err, ErrStyleBorderHorizontalIsUndefined)
	}

	if _, found := s[StyleBorderVertical]; !found {
		return errors.Join(err, ErrStyleBorderVerticalIsUndefined)
	}
	return nil
}

func New(cellWidth uint, addRowDiv bool, dest *os.File) *AsciiTable {

	t := &AsciiTable{
		cellWidth:        cellWidth,
		addRowDiv:        addRowDiv,
		dest:             dest,
		defaultPadding:   PadRight,
		innerCellSpacing: DefaultInnerCellSpacing,
	}
	t.SetTheme(Basic)
	return t
}
