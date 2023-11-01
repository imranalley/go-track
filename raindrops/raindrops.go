package raindrops

import "fmt"

func Convert(number int) string {
	str := ""
	switch number {
	case 3 % number:
		str = fmt.Sprintf("%sPling", str)
	case 5 % number:
		str = fmt.Sprintf("%sPlang", str)
	case 7 % number:
		str = fmt.Sprintf("%sPlong", str)
	}
	if str == "" {
		return fmt.Sprint(number)
	}
	return str
}
