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
		for i := 0; i < len(runes)-1; i++ {
			if runes[i]+1 != runes[i+1] {
				flag = false
				break
			}
		}

		if flag {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
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
