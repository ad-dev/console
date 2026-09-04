package textstyle

import (
	"strconv"
	"unicode/utf8"

	"github.com/ad-dev/console"
)

func (t *textInStyle) SetTruncateIndicator(i string) {
	t.ti = i
}
func (t *textInStyle) Truncate(l int) string {
	if l > -1 {
		t.len = l
	}
	return formatString(t.text, l-utf8.RuneCountInString(t.ti), t.codes...)
}

func (t *textInStyle) String() string {
	return t.Truncate(-1)
}

func (t *textInStyle) Len() int {
	return t.len
}

func FormatString(text string, codes ...byte) string {
	return formatString(text, -1, codes...)
}

func formatString(text string, l int, codes ...byte) string {
	if codes == nil {
		return text
	}
	bb := make([]byte, utf8.RuneCountInString(text)+len(codes))
	bb = append(bb[:0], []byte{ESC, '['}...)
	for _, code := range codes {
		bb = strconv.AppendUint(bb, uint64(code), 10)
		bb = append(bb, ';')
	}
	bb[len(bb)-1] = 'm'
	bb = append(bb, []byte(text)...)
	bb = append(bb, []byte{ESC, '[', '0', 'm'}...)
	return string(bb)
}

func NewTextInStyle(text string, codes ...byte) console.TextElement {
	return &textInStyle{
		len:   utf8.RuneCountInString(text),
		text:  text,
		codes: codes,
	}
}
