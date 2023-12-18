package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d\n", &t)

	for ; t > 0; t-- {
		var text string
		fmt.Scanf("%s\n", &text)

		text = sort(text)

		flag := true
		runes := []rune(text)
		for i := 0; i < len(runes)-1 && flag; i++ {
			flag = runes[i]+1 == runes[i+1]
		}

		fmt.Printf("%s\n", func() string {
			if flag {
				return "YES"
			}
			return "NO"
		}())
	}
}

func sort(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}
