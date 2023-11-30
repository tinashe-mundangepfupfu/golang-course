package main

import "fmt"

func main() {
	xi := make([]int, 0, 10)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	xi = append(xi, 1, 2, 3, 4, 5)
	fmt.Println(xi)
	fmt.Println("-------------------")

	xi = append(xi, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
}
