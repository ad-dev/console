package textstyle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyle(t *testing.T) {
	assert.Equal(t, "test", FormatString("test"))
	for i := range 256 {
		assert.Equal(t, fmt.Sprintf("\033[%dmtest\033[0m", i), FormatString("test", byte(i)))
	}
}

func TestRGBColorStyle(t *testing.T) {
	formattedString := FormatString("this text is in color", 38, 2, 255, 240, 200, 48, 2, 120, 110, 100)
	assert.Equal(
		t,
		"\033[38;2;255;240;200;48;2;120;110;100mthis text is in color\033[0m",
		formattedString,
	)
	t.Log(formattedString)
}
