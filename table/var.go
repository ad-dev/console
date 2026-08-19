package table

import "github.com/ad-dev/console"

var (
	commonStyleBasic = console.Style{
		StyleCorner:            DefaultStyleCorner,
		StyleCornerRight:       DefaultStyleCornerRight,
		StyleCornerBottom:      DefaultStyleCornerBottom,
		StyleCornerBottomRight: DefaultStyleCornerBottomRight,
		StyleCornerJointRight:  DefaultStyleCornerJointRight,
		StyleBorderHorizontal:  DefaultStyleBorderHorizontal,
		StyleBorderVertical:    DefaultStyleBorderVertical,
		StyleBorderJoint:       DefaultStyleBorderJoint,
		StyleBorderJointLeft:   DefaultStyleBorderJointLeft,
		StyleBorderJointRight:  DefaultStyleBorderJointRight,
		StyleBorderJointTop:    DefaultStyleBorderJointTop,
		StyleBorderJointBottom: DefaultStyleBorderJointBottom,
	}

	commonStyleSmooth = console.Style{
		StyleCorner:            "┌",
		StyleCornerRight:       "┐",
		StyleCornerBottom:      "└",
		StyleCornerBottomRight: "┘",
		StyleCornerJointRight:  "┤",
		StyleBorderHorizontal:  "─",
		StyleBorderVertical:    "│",
		StyleBorderJoint:       "┼",
		StyleBorderJointLeft:   "├",
		StyleBorderJointRight:  "┤",
		StyleBorderJointBottom: "┴",
		StyleBorderJointTop:    "┬",
	}

	commonStyleBorderless = console.Style{
		StyleCorner:            " ",
		StyleCornerRight:       " ",
		StyleCornerBottom:      " ",
		StyleCornerBottomRight: " ",
		StyleCornerJointRight:  " ",
		StyleBorderHorizontal:  " ",
		StyleBorderVertical:    " ",
		StyleBorderJoint:       " ",
		StyleBorderJointLeft:   " ",
		StyleBorderJointRight:  " ",
		StyleBorderJointBottom: " ",
		StyleBorderJointTop:    " ",
	}

	commonStyleFancy = console.Style{
		StyleCorner:            " ",
		StyleCornerRight:       " ",
		StyleCornerBottom:      " ",
		StyleCornerBottomRight: " ",
		StyleCornerJointRight:  " ",
		StyleBorderHorizontal:  "=",
		StyleBorderVertical:    " ",
		StyleBorderJoint:       " ",
		StyleBorderJointLeft:   " ",
		StyleBorderJointRight:  " ",
		StyleBorderJointBottom: " ",
		StyleBorderJointTop:    " ",
	}

	commonStyleOldSchool = console.Style{
		StyleCorner:            "╔",
		StyleCornerRight:       "╗",
		StyleCornerBottom:      "╚",
		StyleCornerBottomRight: "╝",
		StyleCornerJointRight:  "╣",
		StyleBorderHorizontal:  "═",
		StyleBorderVertical:    "║",
		StyleBorderJoint:       "╬",
		StyleBorderJointLeft:   "╠",
		StyleBorderJointRight:  "╣",
		StyleBorderJointBottom: "╩",
		StyleBorderJointTop:    "╦",
	}
)

var themeNames = map[console.Theme]string{
	Basic:      "ASCII",
	Smooth:     "Smooth",
	Borderless: "Borderless",
	Fancy:      "Fancy",
	OldSchool:  "Old School",
}

var (
	themes = map[console.Theme]map[console.Section]console.Style{
		Basic: {
			Header: commonStyleBasic,
			Body:   commonStyleBasic,
			Footer: commonStyleBasic,
		},
		Smooth: {
			Header: commonStyleSmooth,
			Body:   commonStyleSmooth,
			Footer: commonStyleSmooth,
		},
		Borderless: {
			Header: commonStyleBorderless,
			Body:   commonStyleBorderless,
			Footer: commonStyleBorderless,
		},
		Fancy: {
			Header: commonStyleFancy,
			Body:   commonStyleFancy,
			Footer: commonStyleFancy,
		},
		OldSchool: {
			Header: commonStyleOldSchool,
			Body:   commonStyleOldSchool,
			Footer: commonStyleOldSchool,
		},
	}
)
