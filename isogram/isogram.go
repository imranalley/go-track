package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	counts := make(map[string]int)
	for _, char := range word {
		if counts[strings.ToLower(string(char))] == 1 && char != ' ' && char != '-' {
			return false
		}
		counts[strings.ToLower(string(char))]++
	}
	return true
}
