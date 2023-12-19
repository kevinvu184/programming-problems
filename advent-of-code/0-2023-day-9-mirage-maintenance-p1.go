package main

import "fmt"

const (
	t = 200
	s = 21
)

func main() {
	sum := 0

	for t := t; t > 0; t-- {
		var input [s]int
		for i := range input {
			fmt.Scanf("%d", &input[i])
		}

		last := make([]int, s)
		a := 0

		last[a] = input[s-1]
		a++

		for l := s - 1; l > 0; l-- {
			extrapolate := make([]int, s)
			for i := 0; i < l; i++ {
				extrapolate[i] = input[i+1] - input[i]
				input[i] = extrapolate[i]
			}

			last[a] = extrapolate[l-1]
			a++
		}

		next := 0
		for _, v := range last {
			next += v
		}
		sum += next
	}

	fmt.Printf("%d\n", sum)
}
