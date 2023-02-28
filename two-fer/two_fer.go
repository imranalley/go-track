package twofer

import "fmt"

// ShareWith takes in a name and returns a string with a two-fer statement
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
