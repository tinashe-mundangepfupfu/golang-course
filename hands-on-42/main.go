package main

import "fmt"

func main() {
	ai := [5]int{} // Composite Type. Array that holds 5 values

	for i := 0; i < len(ai); i++ {
		ai[i] = i
	}

	for i, v := range ai {
		fmt.Printf("type: %T - value %v \n", i, v)
	}
}
