package table

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {

	r, w, e := os.Pipe()
	assert.Nil(t, e)

	tb := New(8, false, w)
	tb.SetTruncateAllCells(true) // enable cells truncation
	tb.SetCellTruncate(2, false) // but disable it for 3rd column
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
		"+---------+---------+-----------------+\n" +
		"|      h1 |      h2 |                 |\n" +
		"+---------+---------+-----------------+\n" +
		"|       1 |       2 |               3 |\n" +
		"+---------+---------+-----------------+\n" +
		"|         |         |Total: something |\n" +
		"+---------+---------+-----------------+\n"

	assert.Equal(t, expectedStr, str)

}

func TestTableWithTruncatedCells(t *testing.T) {

	r, w, e := os.Pipe()
	assert.Nil(t, e)

	tb := New(8, false, w)
	tb.SetTruncateAllCells(true)
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
		"+---------+---------+-----------------+\n" +
		"|      h1 |      h2 |                 |\n" +
		"+---------+---------+-----------------+\n" +
		"|       1 |       2 |               3 |\n" +
		"+---------+---------+-----------------+\n" +
		"|         |         |       Total:... |\n" +
		"+---------+---------+-----------------+\n"

	assert.Equal(t, expectedStr, str)
}

func TestTableWithTruncatedCellsOfLastColumn(t *testing.T) {

	r, w, e := os.Pipe()
	assert.Nil(t, e)

	tb := New(8, false, w)
	tb.SetCellTruncate(2, true)
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
		"+---------+---------+-----------------+\n" +
		"|      h1 |      h2 |                 |\n" +
		"+---------+---------+-----------------+\n" +
		"|       1 |       2 |               3 |\n" +
		"+---------+---------+-----------------+\n" +
		"|         |         |       Total:... |\n" +
		"+---------+---------+-----------------+\n"

	assert.Equal(t, expectedStr, str)
}

func TestGetJSON(t *testing.T) {

	r, w, e := os.Pipe()
	assert.Nil(t, e)

	tb := New(8, false, w)
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
