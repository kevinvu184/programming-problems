package main

import "fmt"

// Time: O(n)
// Space: O(1)
func canConstruct(ransomNote string, magazine string) bool {
	runesNote := []rune(ransomNote)
	m1 := make(map[rune]int)
	for _, v := range runesNote {
		m1[v] += 1
	}

	m2 := make(map[rune]int)
	runesMagazine := []rune(magazine)
	for _, v := range runesMagazine {
		m2[v] += 1
	}

	flag := true
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 > v2 {
			flag = false
			break
		}
	}

	if flag {
		return true
	}
	return false
}

func main() {
	testCases := []struct {
		input struct {
			param1 string
			param2 string
		}
		expected bool
	}{
		{
			input: struct {
				param1 string
				param2 string
			}{
				"a",
				"b",
			},
			expected: false,
		},
		{
			input: struct {
				param1 string
				param2 string
			}{
				"aa",
				"ab",
			},
			expected: false,
		},
		{
			input: struct {
				param1 string
				param2 string
			}{
				"aa",
				"aab",
			},
			expected: true,
		},
	}
	for _, tc := range testCases {
		result := canConstruct(tc.input.param1, tc.input.param2)
		if result != tc.expected {
			fmt.Printf("Expected %v, but got %v\n", tc.expected, result)
		}
	}
}
