package table

var (
	commonStyleBasic = Style{
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

	commonStyleSmooth = Style{
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

	commonStyleBorderless = Style{
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
)

var themeNames = map[Theme]string{
	Basic:      "ASCII",
	Smooth:     "Smooth",
	Borderless: "Borderless",
}

var (
	themes = map[Theme]map[Section]Style{
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
	}
)
