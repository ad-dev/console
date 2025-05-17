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
