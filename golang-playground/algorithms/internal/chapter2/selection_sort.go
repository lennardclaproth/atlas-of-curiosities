package chapter2

import (
	"golang.org/x/exp/constraints"
)

// During each iteration we'll select the smallest item from the
// unsorted partition and move it to the sorted partition
// Track the current minimum and the current item. After each
// Remember to always reset the current minimum to the first
// value of the unsorted partition after which we progress the
// loop until we find a smaller number, swap the elements.
func SelectionSort[T constraints.Ordered](arr []T) {
	n := len(arr)

	// Loop over full array (becomes sorted partition and grows)
	for i := 0; i < n-1; i++ {
		// We need to track the index of the smallest value
		minI := i

		// Loop over unsorted partition and get the smallest 
		// value and its index
		for j := i + 1; j < n; j++ {
			//Check values
			if arr[j] < arr[minI] {
				minI = j
			}
		}

		// Swap elements
		arr[i], arr[minI] = arr[minI], arr[i]
	}
}