package logs

import "unicode/utf8"

func Application(log string) string {
	apps := map[rune]string{
		'\u2757':     "recommendation",
		'\U0001f50d': "search",
		'\u2600':     "weather",
	}
	for _, v := range log {
		if r, name := apps[v]; name {
			return r
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)      // create a slice of runes from log
	for i, v := range runes { // loop thru range of slice runes
		if v == oldRune {
			runes[i] = newRune // replace rune at current index with newRune
		}
	}
	return string(runes) // return runes converted to string
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
