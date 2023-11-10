package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.Trim(id, " ")
	doubledStr := ""
	sum := 0

	if _, err := strconv.Atoi(id); err == nil {
		if len(id) <= 1 {
			return false
		}
		runes := []rune(id)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		for index, n := range runes {
			if index%2 == 0 {
				n *= 2
			}
			doubledStr += string(n)
		}
		for _, s := range doubledStr {
			number, _ := strconv.Atoi(string(s))
			sum += number
		}
	}
	return sum%10 == 0
}
