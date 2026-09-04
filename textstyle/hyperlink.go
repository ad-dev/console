package textstyle

import (
	"unicode/utf8"

	"github.com/ad-dev/console"
)

func (h *hyperlink) SetTruncateIndicator(i string) {
	h.ti = i
}
func (h *hyperlink) Truncate(l int) string {
	if l > -1 {
		h.len = l
	}
	return formatHyperlink(h.url, h.title, l-utf8.RuneCountInString(h.ti))
}

func (h *hyperlink) String() string {
	return h.Truncate(-1)
}

func (h *hyperlink) Len() int {
	return h.len
}

func FormatHyperlink(url, title string) string {
	return formatHyperlink(url, title, -1)
}

func formatHyperlink(url, title string, l int) string {

	bb := make([]byte, len(url)+utf8.RuneCountInString(title))
	bb = append(bb[:0], []byte{ESC, ']', Hyperlink, ';', ';'}...)
	for _, c := range url {
		bb = append(bb, byte(c))
	}
	bb = append(bb, []byte{ESC, '\\'}...)

	if l > 0 {
		title = string([]rune(title)[:l])
	}

	for _, c := range title {
		bb = append(bb, byte(c))
	}
	bb = append(bb, []byte{ESC, ']', Hyperlink, ';', ';', ESC, '\\'}...)
	return string(bb)
}

func NewHyperlink(url, title string) console.TextElement {
	return &hyperlink{
		len:   utf8.RuneCountInString(title),
		url:   url,
		title: title,
	}
}
