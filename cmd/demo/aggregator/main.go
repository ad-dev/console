package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ad-dev/console"
	"github.com/ad-dev/console/table"
	"github.com/ad-dev/console/table/aggregator"
)

func main() {
	t := table.New(20, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2"})
	t.SetTruncateAllCells(true)
	t.AddAnyRow([]any{11, 15})
	t.AddAnyRow([]any{14, 18})
	t.AddAnyRow([]any{05, 56})
	t.AddAnyRow([]any{05, aggregator.SumVertcally().SetPrecision(3)})
	t.AddAnyRow([]any{0.51, 1, aggregator.SumHorizontally().SetPrecision(2)})
	t.AddAnyRow([]any{0.5, 28, 1, aggregator.SumHorizontally().SetPrecision(1)})
	t.AddAnyRow([]any{aggregator.AvgVertcally().SetPrecision(3)})
	t.AddAnyRow([]any{aggregator.MinVertcally().SetPrecision(3)})
	t.AddAnyRow([]any{aggregator.MaxVertcally().SetPrecision(3)})
	t.AddAnyRow([]any{"mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm"})
	t.SetCellFormatter(
		0,
		0,
		func(tc console.TableCell) string {
			str := tc.Value().(string)

			str += "! @" + strconv.Itoa(tc.Row()) + ", " + strconv.Itoa(tc.Column())

			return str
		},
		table.NewCellCondition(0, 0, table.Equals, "11"),
	)
	t.SetTheme(table.Smooth)
	fmt.Println(t.ThemeName())
	t.Display()
}
