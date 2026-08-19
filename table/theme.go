package table

import (
	"fmt"
)

func (t *AsciiTable) ThemeName() string {
	if x, ok := themeNames[t.currentTheme]; ok {
		return x
	}
	return fmt.Sprintf("theme #%d", t.currentTheme)
}
