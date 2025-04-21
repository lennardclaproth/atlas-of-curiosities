package SortingAlgorithms

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{[]int{5, 3, 8, 4, 2}, []int{2, 3, 4, 5, 8}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}}, // Already sorted
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}}, // Reversed
	}

	for _, tt := range tests {
		SelectionSort(tt.input)
		for i, v := range tt.input {
			if v != tt.expected[i] {
				t.Errorf("SelectionSort(%v) = %v; expected %v", tt.input, tt.input, tt.expected)
				break
			}
		}
	}
}
