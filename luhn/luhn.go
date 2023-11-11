package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

func Valid(id string) bool {
	id = strings.Trim(id, " ")
	doubledStr := ""
	sum := 0

	// if _, err := strconv.Atoi(id); err == nil {
	if digitCheck.MatchString(id) {
		if len(id) <= 1 {
			return false
		}
		print(id)
		if len(id)%2 == 0 { // if length is even
			for index, n := range id {
				if index%2 != 0 {
					n *= 2
				}
				doubledStr += string(n)
			}
		} else { // if length is odd
			for index, n := range id {
				if index%2 == 0 {
					n *= 2
				}
				doubledStr += string(n)
			}
		}
		print(doubledStr)
		for _, s := range doubledStr {
			number, _ := strconv.Atoi(string(s))
			sum += number
		}
	}
	return sum%10 == 0
}
