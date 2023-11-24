package resistorcolor

// Colors should return the list of all colors.
func Colors() []string {
	colors := map[]string {
		"black": 0,
		"brown": 1,
		"red": 2,
		"orange": 3,
		"Yellow": 4,
		"Green": 5,
		"blue": 6,
		"violet": 7,
		"grey": 8,
		"white": 9
	}
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	panic("Please implement the ColorCode function")
}
