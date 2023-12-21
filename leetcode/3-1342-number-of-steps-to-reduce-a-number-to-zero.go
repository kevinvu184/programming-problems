package main

import "fmt"

// Time: O(logn)
// Space: O(1)
func numberOfStepsi(num int) int {
	steps := 0
	for ; num != 0; steps++ {
		if num%2 == 0 {
			num /= 2
		} else if num%2 == 1 {
			num -= 1
		}
	}
	return steps
}

// Time: O(logn)
// Space: O(1)
func numberOfStepsr(num int) int {
	if num == 0 {
		return 0
	}
	if num%2 == 0 {
		return numberOfStepsr(num/2) + 1
	}
	return numberOfStepsr(num-1) + 1
}

func main() {
	testCases1 := []struct {
		input    int
		expected int
	}{
		{
			input:    14,
			expected: 6,
		},
		{
			input:    8,
			expected: 4,
		},
		{
			input:    123,
			expected: 12,
		},
	}
	for _, tc := range testCases1 {
		result := numberOfStepsi(tc.input)
		if result != tc.expected {
			fmt.Printf("Expected %v, but got %v\n", tc.expected, result)
		}
	}

	testCases2 := []struct {
		input    int
		expected int
	}{
		{
			input:    14,
			expected: 6,
		},
		{
			input:    8,
			expected: 4,
		},
		{
			input:    123,
			expected: 12,
		},
	}
	for _, tc := range testCases2 {
		result := numberOfStepsr(tc.input)
		if result != tc.expected {
			fmt.Printf("Expected %v, but got %v\n", tc.expected, result)
		}
	}
}
