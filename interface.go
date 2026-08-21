package console

import "os"

type TableCell interface {
	Column() int
	SetRow(int) TableCell
	SetColumn(int) TableCell
	Row() int
	Value() any
	SetValue(any) TableCell
}

type TableRow = []TableCell

type TableCellCondition interface {
	TableCell
	Operator() byte
	SetOperator(byte) TableCellCondition
	Set(row, col int, operator byte, value any) TableCellCondition
}

type TableCellFormatterCallback = func(TableCell) string

type Table interface {
	ClearRows()
	AddHeader(row []string)
	AddAnyHeader(row []any) error
	AddRow(row []string)
	AddAnyRow(row []any) error
	AddFooter(row []string)
	AddAnyFooter(row []any) error
	ChangeRow(index int, row []any) error
	ChangeHeader(row []any) error
	ChangeFooter(row []any) error
	Display() error
	SetDefaultPadding(p byte)
	SetTruncateAllCells(flag bool)
	SetCellTruncate(col uint, flag bool)
	SetColumnPadding(col uint, p byte)
	GetJSON() (string, error)
	GetFormattedJSON() (string, error)
	PrintFormattedJSON(key string) error
	AddCellFormatter(row, col int, cb TableCellFormatterCallback, conditions ...TableCellCondition)
	SetCellFormatter(row, col int, cb TableCellFormatterCallback, conditions ...TableCellCondition)
	SetHeaderStyle(s Style) error
	SetBodyStyle(s Style) error
	SetFooterStyle(s Style) error
	SetTheme(th Theme) error
	SetThemeByName(name string) error
	Theme() Theme
	ThemeName() string
	SetOutputDestination(d *os.File)
}
