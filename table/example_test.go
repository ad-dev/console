package table_test

import (
	"os"

	"github.com/ad-dev/console"
	"github.com/ad-dev/console/table"
)

func Example() {

	t := table.New(8, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2"})
	t.SetDefaultPadding(table.PadLeft)
	t.AddRow([]string{"1", "2", "3\n42\n00"})
	t.AddFooter([]string{"Total: something"})
	t.Display()

	// Output:
	// +----------+----------+------------------+
	// |       h1 |       h2 |                  |
	// +----------+----------+------------------+
	// |        1 |        2 |                3 |
	// |          |          |               42 |
	// |          |          |               00 |
	// +----------+----------+------------------+
	// |          |          | Total: something |
	// +----------+----------+------------------+
}

func ExampleAsciiTable_SetBodyStyle() {

	t := table.New(8, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2"})
	t.SetDefaultPadding(table.PadLeft)
	t.AddRow([]string{"1", "2", "3\n42\n00"})
	t.AddFooter([]string{"Total: something"})
	t.SetBodyStyle(
		console.Style{
			table.StyleCorner:           "+",
			table.StyleBorderHorizontal: " ",
			table.StyleBorderVertical:   ".",
		})
	t.SetFooterStyle(
		console.Style{
			table.StyleCorner:            "_",
			table.StyleCornerRight:       "_",
			table.StyleCornerBottom:      "_",
			table.StyleCornerBottomRight: "_",
			table.StyleCornerJointRight:  "_",

			table.StyleBorderHorizontal:  "_",
			table.StyleBorderVertical:    ":",
			table.StyleBorderJoint:       "_",
			table.StyleBorderJointLeft:   "_",
			table.StyleBorderJointRight:  "_",
			table.StyleBorderJointBottom: "_",
			table.StyleBorderJointTop:    "_",
		})
	t.Display()

	// Output:
	// +----------+----------+------------------+
	// |       h1 |       h2 |                  |
	// +----------+----------+------------------+
	// .        1 .        2 .                3 .
	// .          .          .               42 .
	// .          .          .               00 .
	// __________________________________________
	// :          :          : Total: something :
	// __________________________________________
}
