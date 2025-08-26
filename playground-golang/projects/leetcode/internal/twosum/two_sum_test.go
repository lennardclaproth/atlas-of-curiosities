package twosum

import "testing"

func TestRunCase1CorrectValue(t *testing.T) {
	list := []int{2, 7, 11, 15}
	target := 9

	res := run(list, target)

	if res[0] != 1 && res[1] != 0 {
		t.Errorf("Expected result to be: %v got: %v", []int{0, 1}, res)
	}
}

func TestRunCase2CorrectValue(t *testing.T) {
	list := []int{3, 2, 4}
	target := 6

	res := run(list, target)

	if res[0] != 2 && res[1] != 0 {
		t.Errorf("Expected result to be: %v got: %v", []int{2, 1}, res)
	}
}

func TestRunCase3CorrectValue(t *testing.T) {
	list := []int{3, 3}
	target := 6

	res := run(list, target)

	if res[0] != 1 && res[1] != 0 {
		t.Errorf("Expected result to be: %v got: %v", []int{1, 0}, res)
	}
}
