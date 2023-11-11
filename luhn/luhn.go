package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	sum := 0
	double := len(id)%2 == 0

	// if _, err := strconv.Atoi(id); err == nil {
	if len(id) <= 1 {
		return false
	}
	for _, n := range id {
		if !unicode.IsDigit(n) {
			return false
		}
		num := int(n - '0')
		if double {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
		double = !double
	}
	return sum%10 == 0
}
