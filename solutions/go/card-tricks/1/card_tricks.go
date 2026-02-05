package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if IsIndexInSlice(slice, index) {
		return slice[index]
	}
	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if IsIndexInSlice(slice, index) {
		slice[index] = value
		return slice
	}
	return append(slice, value)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if IsIndexInSlice(slice, index) {
		if index == 0 {
			return slice[1:]
		} else if index+1 == len(slice) {
			return slice[:index]
		} else {
			firstPart := slice[:index]
			secondPart := slice[index+1:]
			return append(firstPart, secondPart...)
		}
	}
	return slice
}

func IsIndexInSlice(slice []int, index int) bool {

	/*for i := range slice {
		if index == i {
			return true
		}
	}
	return false*/
	if index < 0 || index+1 > len(slice) {
		return false
	}
	return true
}
