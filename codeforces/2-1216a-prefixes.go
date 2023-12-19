package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d ", &size)
	var text string
	fmt.Scanf("%s\n", &text)

	runes := []rune(text)
	ops := 0
	for i := 1; i < len(runes); i += 2 {
		if runes[i] == runes[i-1] {
			runes[i] = 'a' + 'b' - runes[i]
			ops++
		}
	}

	fmt.Printf("%d\n%s\n", ops, string(runes))
}
