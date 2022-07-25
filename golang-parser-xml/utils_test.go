package main

import "testing"

func TestFromFloatToCents(t *testing.T) {
	integerNumber := fromFloatToCents(5.05)

	if integerNumber != 505 {
		t.Errorf("Expected to get 505, but got: %v", integerNumber)
	}
}
