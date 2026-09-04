package table

import "github.com/ad-dev/console"

const (
	PadLeft  byte = 2
	PadRight byte = 4

	AlignTop    byte = 8
	AlignBottom byte = 16

	DefaultInnerCellSpacing = 1

	DefaultStyleCorner            = "+"
	DefaultStyleCornerRight       = "+"
	DefaultStyleCornerBottom      = "+"
	DefaultStyleCornerBottomRight = "+"
	DefaultStyleCornerJointRight  = "+"
	DefaultStyleBorderHorizontal  = "-"
	DefaultStyleBorderVertical    = "|"
	DefaultStyleBorderJoint       = "+"
	DefaultStyleBorderJointLeft   = "+"
	DefaultStyleBorderJointRight  = "+"
	DefaultStyleBorderJointTop    = "+"
	DefaultStyleBorderJointBottom = "+"

	StyleCorner            = 1
	StyleCornerRight       = 2
	StyleCornerBottom      = 3
	StyleCornerBottomRight = 4
	StyleCornerJointRight  = 6
	StyleBorderHorizontal  = 7
	StyleBorderVertical    = 8
	StyleBorderJoint       = 9
	StyleBorderJointLeft   = 10
	StyleBorderJointRight  = 11
	StyleBorderJointTop    = 12
	StyleBorderJointBottom = 13
)

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

const (
	HeaderIndex = -10
	FooterIndex = -20
)

const (
	Equals byte = iota
	NotEquals
	LowerThan
	GreaterThan
)

const AnyIndex = -1
