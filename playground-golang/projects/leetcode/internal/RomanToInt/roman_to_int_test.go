package romantoint

import (
	"testing"
)

func TestValueCase1Correct(t *testing.T) {
	res := romanToInt("III")

	if res != 3 {
		t.Errorf("Expected result to be: %v got: %v", 3, res)
	}
}

func TestValueCase2Correct(t *testing.T) {
	res := romanToInt("LVIII")

	if res != 58 {
		t.Errorf("Expected result to be: %v got: %v", 58, res)
	}
}

func TestValueCase3Correct(t *testing.T) {
	res := romanToInt("MCMXCIV")

	if res != 1994 {
		t.Errorf("Expected result to be: %v got: %v", 1994, res)
	}
}
