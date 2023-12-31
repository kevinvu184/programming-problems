package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for ; t > 0; t-- {
		var n int
		fmt.Scanf("%d\n", &n)
		size := 2*n - 2
		inputs := make([]string, size)
		for i := 0; i < size; i++ {
			if i < size-1 {
				fmt.Scanf("%s", &inputs[i])
			} else {
				fmt.Scanf("%s\n", &inputs[i])
			}
		}

		flag := true
		a := ""
		b := ""
		for _, v := range inputs {
			if len(v) == (n-1) && a == "" {
				a = v
			} else if len(v) == (n-1) && a != "" {
				b = v
			}
		}
		x := 0
		y := (n - 1) - 1
		for ; x < (n - 1); x, y = x+1, y-1 {
			if a[x] != b[y] {
				flag = false
			}
		}

		fmt.Printf("%s\n", func() string {
			if flag {
				return "YES"
			}
			return "NO"
		}())
	}
}
