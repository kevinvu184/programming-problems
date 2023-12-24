package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d\n", &size)
	var str string
	fmt.Scanf("%s\n", &str)

	res := 0
	for i, v := range str {
		if rune(v-'0')%2 == 0 {
			res += i + 1
		}
	}

	fmt.Printf("%d\n", res)
}
