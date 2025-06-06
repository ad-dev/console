package table_test

import (
	"os"

	"github.com/ad-dev/console/table"
)

func Example() {

	t := table.New(8, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2"})
	t.AddRow([]string{"1", "2", "3"})
	t.AddFooter([]string{"Total: something"})
	t.Display()

	// Output:
	// +---------+---------+-----------------+
	// |      h1 |      h2 |                 |
	// +---------+---------+-----------------+
	// |       1 |       2 |               3 |
	// +---------+---------+-----------------+
	// |         |         |Total: something |
	// +---------+---------+-----------------+
}
