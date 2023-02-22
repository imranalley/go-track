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

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	_, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}
	number, err := strconv.Atoi(fnb.Value())
	if err != nil {
		return 0
	}
	return number
}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %0.1f", float64(ExtractFancyNumber(fnb)))
}

func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return (DescribeNumber(float64(v)))
	case float64:
		return (DescribeNumber(v))
	case NumberBox:
		return (DescribeNumberBox(v))
	case FancyNumberBox:
		return (DescribeFancyNumberBox(v))
	}
	return "Return to sender"
}
