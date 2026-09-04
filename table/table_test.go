package table

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/ad-dev/console"
	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {

	r, w, e := os.Pipe()
	assert.Nil(t, e)

	tb := New(8, false, w)
	tb.SetTruncateAllCells(true) // enable cells truncation
	tb.SetCellTruncate(2, false) // but disable it for 3rd column
	tb.SetDefaultPadding(PadLeft)
	tb.AddHeader([]string{"h1", "h2"})
	tb.AddRow([]string{"1\n56\nA", "2\nG Y", "3\n42\n00"})
	tb.AddFooter([]string{"Total: something"})
	assert.Nil(t, tb.Display())

	assert.Nil(t, e)

	out := make(chan string)
	go func(t *testing.T) {
		for {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			out <- buf.String()
		}
	}(t)

	w.Close()
	str := <-out
	t.Log("\n\n" + str)

	expectedStr := "" +
		"+----------+----------+------------------+\n" +
		"|       h1 |       h2 |                  |\n" +
		"+----------+----------+------------------+\n" +
		"|        1 |        2 |                3 |\n" +
		"|       56 |      G Y |               42 |\n" +
		"|        A |          |               00 |\n" +
		"+----------+----------+------------------+\n" +
		"|          |          | Total: something |\n" +
		"+----------+----------+------------------+\n"

	assert.Equal(t, expectedStr, str)

}

func TestTableWithTruncatedCells(t *testing.T) {

	r, w, e := os.Pipe()
	assert.Nil(t, e)

	tb := New(8, false, w)
	tb.SetTruncateAllCells(true)
	tb.SetDefaultPadding(PadLeft)
	tb.AddHeader([]string{"h1", "h2"})
	tb.AddRow([]string{"1", "2", "3"})
	tb.AddFooter([]string{"Total: something"})
	assert.Nil(t, tb.Display())

	assert.Nil(t, e)

	out := make(chan string)
	go func(t *testing.T) {
		for {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			out <- buf.String()
		}
	}(t)

	w.Close()
	str := <-out
	t.Log("\n\n" + str)

	expectedStr := "" +
		"+----------+----------+----------+\n" +
		"|       h1 |       h2 |          |\n" +
		"+----------+----------+----------+\n" +
		"|        1 |        2 |        3 |\n" +
		"+----------+----------+----------+\n" +
		"|          |          | Total... |\n" +
		"+----------+----------+----------+\n"

	assert.Equal(t, expectedStr, str)
}

func TestTableWithFormmatters(t *testing.T) {

	r, w, e := os.Pipe()
	assert.Nil(t, e)

	tb := New(8, false, w)
	tb.SetTruncateAllCells(true)
	tb.SetDefaultPadding(PadLeft)
	tb.AddHeader([]string{"h1", "h2"})
	tb.AddRow([]string{"1", "2", "3"})
	tb.AddFooter([]string{"Total: something"})

	cb := func(tc console.TableCell) string {
		return fmt.Sprintf("%s!", tc.Value())
	}

	tb.SetCellFormatter(
		FooterIndex,
		2,
		cb,
		NewCellCondition(0, 0, Equals, "1"),
		NewCellCondition(0, 2, Equals, "3"),
		NewCellCondition(0, 0, LowerThan, "2"),
		NewCellCondition(0, 2, GreaterThan, "2"),
		NewCellCondition(0, 2, NotEquals, "4"),
	)

	assert.Nil(t, tb.Display())

	assert.Nil(t, e)

	out := make(chan string)
	go func(t *testing.T) {
		for {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			out <- buf.String()
		}
	}(t)

	w.Close()
	str := <-out
	t.Log("\n\n" + str)

	expectedStr := "" +
		"+----------+----------+-----------+\n" +
		"|       h1 |       h2 |           |\n" +
		"+----------+----------+-----------+\n" +
		"|        1 |        2 |         3 |\n" +
		"+----------+----------+-----------+\n" +
		"|          |          | Total...! |\n" +
		"+----------+----------+-----------+\n"

	assert.Equal(t, expectedStr, str)
}

func TestTableWithTruncatedCellsOfLastColumn(t *testing.T) {

	r, w, e := os.Pipe()
	assert.Nil(t, e)

	tb := New(8, false, w)
	tb.SetCellTruncate(2, true)
	tb.SetDefaultPadding(PadLeft)
	tb.AddHeader([]string{"h1", "h2"})
	tb.AddRow([]string{"1", "2", "3"})
	tb.AddFooter([]string{"Total: something"})
	assert.Nil(t, tb.Display())

	assert.Nil(t, e)

	out := make(chan string)
	go func(t *testing.T) {
		for {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			out <- buf.String()
		}
	}(t)

	w.Close()
	str := <-out
	t.Log("\n\n" + str)

	expectedStr := "" +
		"+----------+----------+----------+\n" +
		"|       h1 |       h2 |          |\n" +
		"+----------+----------+----------+\n" +
		"|        1 |        2 |        3 |\n" +
		"+----------+----------+----------+\n" +
		"|          |          | Total... |\n" +
		"+----------+----------+----------+\n"

	assert.Equal(t, expectedStr, str)
}

func TestGetJSON(t *testing.T) {

	r, w, e := os.Pipe()
	assert.Nil(t, e)

	tb := New(8, false, w)
	tb.SetDefaultPadding(PadLeft)
	tb.AddHeader([]string{"h1", "h2"})
	tb.AddRow([]string{"1", "2", "3"})
	tb.AddFooter([]string{"Total: something"})

	assert.Nil(t, e)

	out := make(chan string)
	go func(t *testing.T) {
		for {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			out <- buf.String()
		}
	}(t)

	w.Close()
	str := <-out
	t.Log("\n\n" + str)

	jsonStr, err := tb.GetJSON()
	assert.Nil(t, err)
	assert.Equal(t, `[["h1","h2",""],["1","2","3"],["","","Total: something"]]`, jsonStr)

}
