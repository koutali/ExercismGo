package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("%s%.1f", "This is the number ", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("%s%.1f", "This is a box containing the number ", float64(nb.Number()))
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
	fancyNumber, isFancyNumber := fnb.(FancyNumber)
	if !isFancyNumber {
		return 0
	}

	fancyNumberValue, isError := strconv.Atoi(fancyNumber.Value())
	if isError == nil {
		return fancyNumberValue
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("%s%.1f", "This is a fancy box containing the number ", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	case NumberBox:
		return DescribeNumberBox(v)
	default:
		return "Return to sender"
	}

	return "Return to sender"
}
