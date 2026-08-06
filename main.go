package main

// live progress bar inside table demo

import (
	"fmt"
	"os"
	"time"

	"github.com/ad-dev/console/progressbar"
	"github.com/ad-dev/console/table"
)

func main() {
	pb := progressbar.New(0, 10)
	pb2 := progressbar.New(0, 5)
	t := table.New(20, false, os.Stdout)
	t.AddHeader([]string{"h1", "h2"})
	t.SetDefaultPadding(table.PadLeft)

	t.AddRow([]string{"", "2", "3\n42\n00"})
	t.AddRow([]string{pb.String(), "9", "3\n02\n15"})
	t.AddFooter([]string{"Total: something", pb2.String()})

	for !pb.IsComplete() || !pb2.IsComplete() {
		t.ChangeRow(1, []any{pb.String(), "9", "3\n02\n15"})
		t.ChangeFooter([]any{"Total: something", pb2.String()})
		t.Display()
		fmt.Printf("\033[12F\033[m")
		pb.Inc(5)
		pb2.Inc(15.123)
		time.Sleep(500 * time.Millisecond)
	}
	t.ChangeRow(1, []any{pb.String(), "9", "3\n02\n15"})
	t.ChangeFooter([]any{"Total: something", pb2.String()})
	t.Display()
}
