package example_test

import (
	"fmt"
	"os"

	"github.com/ad-dev/console/progressbar"
	"github.com/ad-dev/console/table"
)

func Example() {

	t := table.New(8, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2"})
	t.SetCellTruncate(0, true)
	t.SetDefaultPadding(table.PadLeft)
	t.AddRow([]string{"", "2", "3\n42\n00"})
	t.AddRow([]string{progressbar.New(0, 5).String(), "9", "3\n02\n15"})

	t.AddFooter([]string{"Total: something", progressbar.New(65, 5).String()})
	t.Display()

	// Output:
	// +--------------------+------------------+-------------------------+
	// |                 h1 |               h2 |                         |
	// +--------------------+------------------+-------------------------+
	// |                    |                2 |                       3 |
	// |                    |                  |                      42 |
	// |                    |                  |                      00 |
	// |           ░░░░░... |                9 |                       3 |
	// |                    |                  |                      02 |
	// |                    |                  |                      15 |
	// +--------------------+------------------+-------------------------+
	// |                    | Total: something |           ███▒░  65.00% |
	// +--------------------+------------------+-------------------------+
}

func ExampleAsciiTable_SetTheme() {

	t := table.New(6, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2", "long header title lllllllllllllllllll"})
	t.SetDefaultPadding(table.PadLeft)
	t.AddRow([]string{"1", "2", "3\n42\n00"})
	t.AddAnyRow([]any{table.CustomCellWidth{Content: "abc", Width: 30}})
	t.AddFooter([]string{"Total: something"})

	t.SetTheme(table.Basic)
	fmt.Printf("\nTheme name: %s\n", t.ThemeName())
	t.Display()

	t.SetTheme(table.Smooth)
	fmt.Printf("\nTheme name: %s\n", t.ThemeName())
	t.Display()

	t.SetTheme(table.OldSchool)
	fmt.Printf("\nTheme name: %s\n", t.ThemeName())
	t.Display()
	// Output:
	//
	// Theme name: ASCII
	// +--------------------------------+--------+---------------------------------------+
	// |                             h1 |     h2 | long header title lllllllllllllllllll |
	// +--------------------------------+--------+---------------------------------------+
	// |                              1 |      2 |                                     3 |
	// |                                |        |                                    42 |
	// |                                |        |                                    00 |
	// |                            abc |        |                                       |
	// +--------------------------------+--------+---------------------------------------+
	// |                                |        |                      Total: something |
	// +--------------------------------+--------+---------------------------------------+
	//
	// Theme name: Smooth
	// ┌────────────────────────────────┬────────┬───────────────────────────────────────┐
	// │                             h1 │     h2 │ long header title lllllllllllllllllll │
	// ├────────────────────────────────┼────────┼───────────────────────────────────────┤
	// │                              1 │      2 │                                     3 │
	// │                                │        │                                    42 │
	// │                                │        │                                    00 │
	// │                            abc │        │                                       │
	// ├────────────────────────────────┼────────┼───────────────────────────────────────┤
	// │                                │        │                      Total: something │
	// └────────────────────────────────┴────────┴───────────────────────────────────────┘
	//
	// Theme name: Old School
	// ╔════════════════════════════════╦════════╦═══════════════════════════════════════╗
	// ║                             h1 ║     h2 ║ long header title lllllllllllllllllll ║
	// ╠════════════════════════════════╬════════╬═══════════════════════════════════════╣
	// ║                              1 ║      2 ║                                     3 ║
	// ║                                ║        ║                                    42 ║
	// ║                                ║        ║                                    00 ║
	// ║                            abc ║        ║                                       ║
	// ╠════════════════════════════════╬════════╬═══════════════════════════════════════╣
	// ║                                ║        ║                      Total: something ║
	// ╚════════════════════════════════╩════════╩═══════════════════════════════════════╝
}
