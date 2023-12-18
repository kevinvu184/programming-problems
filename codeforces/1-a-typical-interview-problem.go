package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)

	for ; t > 0; t-- {
		var nums int
		var str string
		fmt.Scanf("%d\n%s\n", &nums, &str)

		flag := false

		sample := "FBFFBFFBFBFFBFFBFBF"
		for i := 0; i+nums < len(sample) && !flag; i++ {
			substr := sample[i : i+nums]
			if substr == str {
				flag = true
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
