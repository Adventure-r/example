package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
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

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	number, isFancyNumber := fnb.(FancyNumber)
	if isFancyNumber {
		v, _ := strconv.Atoi(number.Value())
		return v
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	number, isFancyNumber := fnb.(FancyNumber)
	if isFancyNumber {
		v, ok := strconv.Atoi(number.Value())
		if ok == nil {
			return fmt.Sprintf("This is a fancy box containing the number %d.0", v)
		}
	}
	return "This is a fancy box containing the number 0.0"

}

// DescribeAnything should return a string describing whatever it contains.
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
	}
	return "Return to sender"
}
