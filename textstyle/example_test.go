package textstyle_test

import (
	"fmt"

	"github.com/ad-dev/console/textstyle"
)

func Example() {
	fmt.Println(
		textstyle.FormatString("this text is in color", 38, 2, 255, 240, 200, 48, 2, 120, 110, 100),
	)
}
