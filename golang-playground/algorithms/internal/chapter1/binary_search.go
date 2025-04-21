package chapter1

import "golang.org/x/exp/constraints"

// Assuming the arr is sorted the binarysearch goes as follows,
// it splits the array in two equal parts and checks if the value
// should fall in the left half or the right half. It does so by
// checking what the value is of the middle. This can be done
// recursively.
func BinarySearch[T constraints.Ordered](arr []T, val T) int {
	// If the length of the array is 0, then the value is not
	// in the array
	if len(arr) == 0 {
		return -1
	}

	// Find the midpoint of the array
	mid := len(arr) / 2

	// If the value is equal to the midpoint value
	// return the midpoint value
	if val == arr[mid] {
		return mid
	}

	// If the value is smaller then the midpoint
	// value search left part
	if val < arr[mid] {
		return BinarySearch(arr[:mid], val)
	}

	// search right part
	rightIndex := BinarySearch(arr[mid+1:], val)
	if rightIndex == -1 {
		return -1
	}
	return mid + 1 + rightIndex
}
