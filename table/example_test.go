package table_test

import (
	"os"

	"github.com/ad-dev/console/table"
)

func Example() {

	t := table.New(8, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2"})
	t.AddRow([]string{"1", "2", "3\n42\n00"})
	t.AddFooter([]string{"Total: something"})
	t.Display()

	// Output:
	// +---------+---------+-----------------+
	// |      h1 |      h2 |                 |
	// +---------+---------+-----------------+
	// |       1 |       2 |               3 |
	// |         |         |              42 |
	// |         |         |              00 |
	// +---------+---------+-----------------+
	// |         |         |Total: something |
	// +---------+---------+-----------------+
}

func ExampleAsciiTable_SetBodyStyle() {

	t := table.New(8, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2"})
	t.AddRow([]string{"1", "2", "3\n42\n00"})
	t.AddFooter([]string{"Total: something"})
	t.SetBodyStyle(
		table.Style{
			table.STYLE_CORNER:            "+",
			table.STYLE_BORDER_HORIZONTAL: " ",
			table.STYLE_BORDER_VERTICAL:   ".",
		})
	t.SetFooterStyle(
		table.Style{
			table.STYLE_CORNER:              "_",
			table.STYLE_CORNER_RIGHT:        "_",
			table.STYLE_CORNER_BOTTOM:       "_",
			table.STYLE_CORNER_BOTTOM_RIGHT: "_",
			table.STYLE_CORNER_JOINT_RIGHT:  "_",

			table.STYLE_BORDER_HORIZONTAL:   "_",
			table.STYLE_BORDER_VERTICAL:     ":",
			table.STYLE_BORDER_JOINT:        "_",
			table.STYLE_BORDER_JOINT_LEFT:   "_",
			table.STYLE_BORDER_JOINT_RIGHT:  "_",
			table.STYLE_BORDER_JOINT_BOTTOM: "_",
			table.STYLE_BORDER_JOINT_TOP:    "_",
		})
	t.Display()

	// Output:
	// +---------+---------+-----------------+
	// |      h1 |      h2 |                 |
	// +---------+---------+-----------------+
	// .       1 .       2 .               3 .
	// .         .         .              42 .
	// .         .         .              00 .
	// _______________________________________
	// :         :         :Total: something :
	// _______________________________________
}
