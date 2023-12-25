package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for ; t > 0; t-- {
		var s string
		fmt.Scanf("%s\n", &s)
		r1 := []rune(s)
		var t string
		fmt.Scanf("%s\n", &t)

		for i := 0; i < len(r1); i++ {
			for j := i + 1; j < len(r1); j++ {
				if r1[i] > r1[j] {
					r1[i], r1[j] = r1[j], r1[i]
				}
			}
		}

		a := 0
		b := 0
		c := 0
		d := len(r1)
		for i := 0; i < len(r1); i++ {
			if r1[i] == 'a' {
				a++
			} else if r1[i] == 'b' {
				b++
			} else if r1[i] == 'c' {
				c++
			} else if d == len(r1) {
				d = i
			}
		}

		if a == 0 || b == 0 || c == 0 || t != "abc" {
			fmt.Printf("%s\n", string(r1))
			continue
		}

		for i := 0; i < a; i++ {
			fmt.Printf("%c", 'a')
		}
		for i := 0; i < c; i++ {
			fmt.Printf("%c", 'c')
		}
		for i := 0; i < b; i++ {
			fmt.Printf("%c", 'b')
		}
		fmt.Printf("%s\n", string(r1[d:]))
	}
}
