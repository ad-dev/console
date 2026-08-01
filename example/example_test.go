package example_test

import (
	"fmt"
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

func ExampleAsciiTable_SetTheme() {

	t := table.New(8, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2", "long header title lllllllllllllllllll"})
	t.AddRow([]string{"1", "2", "3\n42\n00"})
	t.AddAnyRow([]any{table.CustomCellWidth{Content: "abc", Width: 30}})
	t.AddFooter([]string{"Total: something"})

	t.SetTheme(table.Basic)
	fmt.Printf("\nTheme name: %s\n", t.Theme())
	t.Display()

	t.SetTheme(table.Smooth)
	fmt.Printf("\nTheme name: %s\n", t.Theme())
	t.Display()
	// Output:
	//
	// Theme name: ASCII
	// +-------------------------------+---------+--------------------------------------+
	// |                            h1 |      h2 |long header title lllllllllllllllllll |
	// +-------------------------------+---------+--------------------------------------+
	// |                             1 |       2 |                                    3 |
	// |                               |         |                                   42 |
	// |                               |         |                                   00 |
	// |                           abc |         |                                      |
	// +-------------------------------+---------+--------------------------------------+
	// |                               |         |                     Total: something |
	// +-------------------------------+---------+--------------------------------------+
	//
	// Theme name: Smooth
	// ┌───────────────────────────────┬─────────┬──────────────────────────────────────┐
	// │                            h1 │      h2 │long header title lllllllllllllllllll │
	// ├───────────────────────────────┼─────────┼──────────────────────────────────────┤
	// │                             1 │       2 │                                    3 │
	// │                               │         │                                   42 │
	// │                               │         │                                   00 │
	// │                           abc │         │                                      │
	// ├───────────────────────────────┼─────────┼──────────────────────────────────────┤
	// │                               │         │                     Total: something │
	// └───────────────────────────────┴─────────┴──────────────────────────────────────┘
}
