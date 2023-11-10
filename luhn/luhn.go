package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.Trim(id, " ")
	doubledStr := ""
	sum := 0

	fmt.Print(len(id))
	if len(id) <= 1 {
		return false
	}
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
	fmt.Printf("Number: %s", id)
	for _, s := range doubledStr {
		number, _ := strconv.Atoi(string(s))
		sum += number
	}
	return sum%10 == 0
}
