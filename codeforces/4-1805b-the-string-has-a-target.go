package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for ; t > 0; t-- {
		var size int
		fmt.Scanf("%d\n", &size)
		var txt string
		fmt.Scanf("%s\n", &txt)

		runes := []rune(txt)
		l, li := 'z', 0
		for i, v := range runes {
			if v <= l {
				l, li = v, i
			}
		}
		if l <= runes[0] {
			for i := li; i > 0; i-- {
				runes[i] = runes[i-1]
			}
			runes[0] = l
		}
		fmt.Printf("%s\n", string(runes))
	}
}
