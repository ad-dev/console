package main

import (
	"fmt"
	"os"

	"github.com/ad-dev/console/table"
	"github.com/ad-dev/console/table/aggregator"
)

func main() {
	t := table.New(20, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2"})
	t.AddAnyRow([]any{11, 15})
	t.AddAnyRow([]any{14, 18})
	t.AddAnyRow([]any{05, 56})
	t.AddAnyRow([]any{05, aggregator.SumVertcally().SetPrecision(3)})
	t.AddAnyRow([]any{0.51, 1, aggregator.SumHorizontally().SetPrecision(2)})
	t.AddAnyRow([]any{0.5, 28, 1, aggregator.SumHorizontally().SetPrecision(1)})

	t.SetTheme(table.Smooth)
	fmt.Println(t.ThemeName())
	t.Display()
}
