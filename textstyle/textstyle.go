package textstyle

import (
	"strconv"
)

func FormatString(text string, codes ...byte) string {
	if codes == nil {
		return text
	}
	bb := make([]byte, len(text)+len(codes))
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
