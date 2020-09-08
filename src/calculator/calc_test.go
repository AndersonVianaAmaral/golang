package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	var result = Sum(5, 5)

	if result != 10 {
		t.Errorf("Sum() is Invalid: returned %v, expected %v ", result, 10)
	}
}
