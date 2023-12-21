package main

import (
	"fmt"
	"slices"
	"strconv"
)

// Time: O(n)
// Space: O(1)
func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 {
			result[i] += "Fizz"
		}
		if (i+1)%5 == 0 {
			result[i] += "Buzz"
		}
		if result[i] == "" {
			result[i] = strconv.Itoa(i + 1)
		}
	}
	return result
}

func main() {
	testCases := []struct {
		input    int
		expected []string
	}{
		{
			input:    3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			input:    5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			input:    15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, tc := range testCases {
		result := fizzBuzz(tc.input)
		if !slices.Equal(result, tc.expected) {
			fmt.Printf("Expected %v, but got %v\n", tc.expected, result)
		}
	}
}
