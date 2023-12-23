package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d\n", &t)

	for ; t > 0; t-- {
		var size int
		fmt.Scanf("%d\n", &size)
		var plays string
		fmt.Scanf("%s\n", &plays)

		fmt.Printf("%c\n", []rune(plays)[size-1])
	}
}
