package textstyle

type hyperlink struct {
	len   int
	ti    string
	url   string
	title string
}

type textInStyle struct {
	len   int
	ti    string
	text  string
	codes []byte
}
