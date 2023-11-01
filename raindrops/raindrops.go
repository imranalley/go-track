package raindrops

import "fmt"

func Convert(number int) string {
	str := ""
	if number%3 == 0 {
		str = fmt.Sprintf("%sPling", str)
	}
	if number%5 == 0 {
		str = fmt.Sprintf("%sPlang", str)
	}
	if number%7 == 0 {
		str = fmt.Sprintf("%sPlong", str)
	}
	if str == "" {
		return fmt.Sprint(number)
	}
	return str
}
