package table

import "fmt"

func (t Theme) String() string {
	if x, ok := themeNames[t]; ok {
		return x
	}
	return fmt.Sprintf("theme #%d", t)
}
