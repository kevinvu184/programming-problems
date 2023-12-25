package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d\n", &size)

	aabb := "aabb"
	for i := 0; i < size; i++ {
		fmt.Printf("%c", aabb[i%4])
	}
	fmt.Printf("\n")
}
