package table

import "github.com/ad-dev/console"

const (
	Header console.Section = iota
	Body
	Footer
)

const (
	Borderless console.Theme = iota
	Basic
	Smooth
	Fancy
	OldSchool
)
