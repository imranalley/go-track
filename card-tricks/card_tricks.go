package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	// panic("Please implement the FavoriteCards function")
	favCards := []int{2,6,9}
	return (favCards[:])
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	// panic("Please implement the GetItem function")
	if(index >= 0 && index < cap(slice)){
		return (slice[index])
	}
	return(-1)
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	// panic("Please implement the SetItem function")
	if(index >= 0 && index < cap(slice)){
		slice [index] = value
		return slice[:]
	}
	slice = append(slice, value)
	return slice[:]
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	// panic("Please implement the PrependItems function")
	if (len(values) !=0 ) {
		slice = append(values, slice...)
		return slice[:]
	}
	return slice[:]
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	// panic("Please implement the RemoveItem function")
	if index >= 0 && index < cap(slice) {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice [:]
}
