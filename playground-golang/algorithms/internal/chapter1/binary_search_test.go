package chapter1

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	// Test cases for integers
	intTests := []struct {
		arr      []int
		val      int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},       // 3 is at index 2
		{[]int{1, 2, 3, 4, 5}, 1, 0},       // 1 is at index 0
		{[]int{1, 2, 3, 4, 5}, 5, 4},       // 5 is at index 4
		{[]int{10, 20, 30, 40, 50}, 40, 3}, // 40 is at index 3
		{[]int{5, 6, 7, 8, 9}, 6, 1},       // 6 is at index 1
		{[]int{1, 2, 3, 4, 5}, 10, -1},     // 10 is not in the array
	}

	for _, tt := range intTests {
		got := BinarySearch(tt.arr, tt.val)
		if got != tt.expected {
			t.Errorf("BinarySearch(%v, %d) = %d; expected %d", tt.arr, tt.val, got, tt.expected)
		}
	}

	// Test cases for floats
	floatTests := []struct {
		arr      []float64
		val      float64
		expected int
	}{
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 3.3, 2},
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 1.1, 0},
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 5.5, 4},
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 6.6, -1}, // value not present
	}

	for _, tt := range floatTests {
		got := BinarySearch(tt.arr, tt.val)
		if got != tt.expected {
			t.Errorf("BinarySearch(%v, %.1f) = %d; expected %d", tt.arr, tt.val, got, tt.expected)
		}
	}

	// Test cases for strings
	stringTests := []struct {
		arr      []string
		val      string
		expected int
	}{
		{[]string{"apple", "banana", "cherry", "date", "fig"}, "cherry", 2},
		{[]string{"apple", "banana", "cherry", "date", "fig"}, "apple", 0},
		{[]string{"apple", "banana", "cherry", "date", "fig"}, "fig", 4},
		{[]string{"apple", "banana", "cherry", "date", "fig"}, "grape", -1}, // not present
	}

	for _, tt := range stringTests {
		got := BinarySearch(tt.arr, tt.val)
		if got != tt.expected {
			t.Errorf("BinarySearch(%v, %s) = %d; expected %d", tt.arr, tt.val, got, tt.expected)
		}
	}
}
