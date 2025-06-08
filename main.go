package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := []int{3, 80, 5, 9}
	result := LargestNumber(input)
	fmt.Println("Input:", input)
	fmt.Println("Result:", result)
}

// LargestNumber returns the largest number that can be formed by concatenating the input integers.
func LargestNumber(nums []int) string {
	// Convert each int to string
	numStrs := make([]string, len(nums))
	for i, num := range nums {
		numStrs[i] = strconv.Itoa(num)
	}

	// Custom sort: compare by concatenated value
	sort.Slice(numStrs, func(i, j int) bool {
		return numStrs[i]+numStrs[j] > numStrs[j]+numStrs[i]
	})

	// If the first element is "0", then the entire number is "0"
	if numStrs[0] == "0" {
		return "0"
	}

	// Join all strings
	return strings.Join(numStrs, "")
}
