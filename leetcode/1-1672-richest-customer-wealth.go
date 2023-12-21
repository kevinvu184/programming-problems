package main

import "fmt"

// Time: O(n^2)
// Space: O(1)
func maximumWealth(input [][]int) int {
	max := 0
	for _, account := range input {
		total := 0
		for _, money := range account {
			total += money
		}
		if total > max {
			max = total
		}
	}
	return max
}

func main() {
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{1, 2, 3}, {3, 2, 1}},
			expected: 6,
		},
		{
			input:    [][]int{{1, 5}, {7, 3}, {3, 5}},
			expected: 10,
		},
		{
			input:    [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			expected: 17,
		},
	}
	for _, tc := range testCases {
		result := maximumWealth(tc.input)
		if result != tc.expected {
			fmt.Printf("Expected %v, but got %v\n", tc.expected, result)
		}
	}
}
