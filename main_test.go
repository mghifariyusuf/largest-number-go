package main

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected string
	}{
		{[]int{10, 2}, "210"},
		{[]int{3, 80, 5, 9}, "98053"},
		{[]int{0, 0}, "0"},
		{[]int{1, 20, 23, 4, 8}, "8423201"},
		{[]int{121, 12}, "12121"},
	}

	for _, test := range tests {
		result := LargestNumber(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %q but got %q", test.input, test.expected, result)
		}
	}
}
