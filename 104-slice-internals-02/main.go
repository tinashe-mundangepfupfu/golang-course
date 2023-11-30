package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	// b := a
	b := make([]int, len(a)) // make([]T, length, capacity)

	copy(b, a) // copy(dst, src)

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("-------------")

	a[0] = 7
	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("-------------")
}
