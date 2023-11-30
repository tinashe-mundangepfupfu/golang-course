package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(xi)

	fmt.Println("-------------------")

	xi = append(xi, 10)

	fmt.Println(xi)
	fmt.Println("-------------------")
}
