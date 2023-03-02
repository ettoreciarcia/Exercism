package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}

	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	var newSlice []int
	if len(slice) <= index || index < 0 {
		newSlice = slice[:]
		newSlice = append(newSlice, value)
	} else {
		newSlice = slice[:]
		newSlice[index] = value
	}

	return newSlice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	newSlice := make([]int, len(slice)+len(values))
	copy(newSlice[len(values):], slice)
	copy(newSlice, values)

	return newSlice

}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	} else {
		newSlice := make([]int, 0, len(slice)-1)
		newSlice = append(newSlice, slice[:index]...)
		newSlice = append(newSlice, slice[index+1:]...)
		return newSlice
	}

}
