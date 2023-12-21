package main

import (
	"fmt"
	"slices"
)

// Time: O(n)
// Space: O(1)
func runningSum1(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

// Time: O(n)
// Space: O(1)
func runningSum2(nums []int) []int {
	prev := 0
	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		prev += nums[i]
		result[i] = prev
	}
	return result
}

func main() {
	var testCases1 = []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    []int{3, 1, 2, 10, 1},
			expected: []int{3, 4, 6, 16, 17},
		},
	}
	for _, tc := range testCases1 {
		result := runningSum1(tc.input)
		if !slices.Equal(result, tc.expected) {
			fmt.Printf("Expected %v, but got %v\n", tc.expected, result)
		}
	}

	var testCases2 = []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    []int{3, 1, 2, 10, 1},
			expected: []int{3, 4, 6, 16, 17},
		},
	}
	for _, tc := range testCases2 {
		result := runningSum2(tc.input)
		if !slices.Equal(result, tc.expected) {
			fmt.Printf("Expected %v, but got %v\n", tc.expected, result)
		}
	}
}
