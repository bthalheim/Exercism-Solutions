package sorting

import (
	"fmt"
	"strconv"
)

func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)

}

type NumberBox interface {
	Number() int
}

func DescribeNumberBox(nb NumberBox) string {

	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))

}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

func ExtractFancyNumber(fnb FancyNumberBox) int {

	_, ok := fnb.(FancyNumber)

	if ok {

		num, _ := strconv.Atoi(fnb.Value())
		return num

	}

	return 0

}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {

	_, ok := fnb.(FancyNumber)

	if ok {

		num, _ := strconv.Atoi(fnb.Value())

		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(num))

	}

	return "This is a fancy box containing the number 0.0"

}

func DescribeAnything(i any) string {

	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}

}
