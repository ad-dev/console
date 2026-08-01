package table

var (
	commonStyleBasic = Style{
		STYLE_CORNER:              DEFAULT_STYLE_CORNER,
		STYLE_CORNER_RIGHT:        DEFAULT_STYLE_CORNER_RIGHT,
		STYLE_CORNER_BOTTOM:       DEFAULT_STYLE_CORNER_BOTTOM,
		STYLE_CORNER_BOTTOM_RIGHT: DEFAULT_STYLE_CORNER_BOTTOM_RIGHT,
		STYLE_CORNER_JOINT_RIGHT:  DEFAULT_STYLE_CORNER_JOINT_RIGHT,
		STYLE_BORDER_HORIZONTAL:   DEFAULT_STYLE_BORDER_HORIZONTAL,
		STYLE_BORDER_VERTICAL:     DEFAULT_STYLE_BORDER_VERTICAL,
		STYLE_BORDER_JOINT:        DEFAULT_STYLE_BORDER_JOINT,
		STYLE_BORDER_JOINT_LEFT:   DEFAULT_STYLE_BORDER_JOINT_LEFT,
		STYLE_BORDER_JOINT_RIGHT:  DEFAULT_STYLE_BORDER_JOINT_RIGHT,
		STYLE_BORDER_JOINT_TOP:    DEFAULT_STYLE_BORDER_JOINT_TOP,
		STYLE_BORDER_JOINT_BOTTOM: DEFAULT_STYLE_BORDER_JOINT_BOTTOM,
	}

	commonStyleSmooth = Style{
		STYLE_CORNER:              "┌",
		STYLE_CORNER_RIGHT:        "┐",
		STYLE_CORNER_BOTTOM:       "└",
		STYLE_CORNER_BOTTOM_RIGHT: "┘",
		STYLE_CORNER_JOINT_RIGHT:  "┤",
		STYLE_BORDER_HORIZONTAL:   "─",
		STYLE_BORDER_VERTICAL:     "│",
		STYLE_BORDER_JOINT:        "┼",
		STYLE_BORDER_JOINT_LEFT:   "├",
		STYLE_BORDER_JOINT_RIGHT:  "┤",
		STYLE_BORDER_JOINT_BOTTOM: "┴",
		STYLE_BORDER_JOINT_TOP:    "┬",
	}

	commonStyleBorderless = Style{
		STYLE_CORNER:              " ",
		STYLE_CORNER_RIGHT:        " ",
		STYLE_CORNER_BOTTOM:       " ",
		STYLE_CORNER_BOTTOM_RIGHT: " ",
		STYLE_CORNER_JOINT_RIGHT:  " ",
		STYLE_BORDER_HORIZONTAL:   " ",
		STYLE_BORDER_VERTICAL:     " ",
		STYLE_BORDER_JOINT:        " ",
		STYLE_BORDER_JOINT_LEFT:   " ",
		STYLE_BORDER_JOINT_RIGHT:  " ",
		STYLE_BORDER_JOINT_BOTTOM: " ",
		STYLE_BORDER_JOINT_TOP:    " ",
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
