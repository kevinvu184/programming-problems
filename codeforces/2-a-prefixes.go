package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d ", &size)
	var text string
	fmt.Scanf("%s\n", &text)

	runes := []rune(text)
	ops := 0
	for i := 0; i < len(runes); i++ {
		if i%2 == 1 {
			if runes[i] == runes[i-1] {
				if runes[i] == 'a' {
					runes[i] = 'b'
				} else {
					runes[i] = 'a'
				}
				ops++
			}
		}
	}

	fmt.Printf("%d\n%s\n", ops, string(runes))
}
