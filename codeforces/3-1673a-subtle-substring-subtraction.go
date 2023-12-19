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
				} else if len(runes) > 1 {
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

func sum(r []rune) int {
	sum := 0
	for i := 0; i < len(r); i++ {
		sum += int(r[i]-'0') - 48
	}
	return sum
}

// Remove a rune from the slice while maintaining order
// a := []rune{'A', 'B', 'C', 'D', 'E'}
// i := 2
// fmt.Printf("%c %d\n", a, len(a))
// copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
// a[len(a)-1] = 0      // Erase last element (write zero value).
// a = a[:len(a)-1]     // Truncate slice.
// fmt.Printf("%c %d\n", a, len(a))
