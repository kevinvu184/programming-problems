package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	s := make([]string, n)
	for i := range s {
		if i < n-1 {
			fmt.Scanf("%s", &s[i])
		} else {
			fmt.Scanf("%s\n", &s[i])
		}
	}

	t := make([]string, m)
	for i := range t {
		if i < m-1 {
			fmt.Scanf("%s", &t[i])
		} else {
			fmt.Scanf("%s\n", &t[i])
		}
	}

	var q int
	fmt.Scanf("%d\n", &q)

	years := make([]int, q)
	for i := range years {
		fmt.Scanf("%d\n", &years[i])
		fmt.Printf("%s%s\n", s[(years[i]-1)%n], t[(years[i]-1)%m])
	}
}
