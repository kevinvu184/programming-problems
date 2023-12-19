package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d\n", &t)

	for ; t > 0; t-- {
		a := 0
		b := 0

		var str string
		fmt.Scanf("%s\n", &str)
		runes := []rune(str)
		for i := 0; len(runes) > 0; i++ {
			if i%2 == 0 {
				if len(runes)%2 == 0 {
					a += sum(runes)
					runes = []rune{}
				} else {
					if runes[0] > runes[len(runes)-1] {
						a += sum(runes[0 : len(runes)-1])
						runes = runes[len(runes)-1:]
					} else {
						a += sum(runes[1:])
						runes = runes[0:1]
					}
				}
			} else {
				if len(runes)%2 == 0 {
					if runes[0] > runes[len(runes)-1] {
						b += sum(runes[0 : len(runes)-2])
						runes = runes[len(runes)-1:]
					} else {
						b += sum(runes[1 : len(runes)-1])
						runes = runes[0:1]
					}
				} else {
					b += sum(runes)
					runes = []rune{}
				}
			}
		}

		name, point := func() (string, int) {
			if a > b {
				return "Alice", a - b
			}
			return "Bob", b - a
		}()
		fmt.Printf("%s %d\n", name, point)
	}
}

func sum(rs []rune) int {
	sum := 0
	for _, r := range rs {
		sum += int(r-'0') - 48
	}
	return sum
}
