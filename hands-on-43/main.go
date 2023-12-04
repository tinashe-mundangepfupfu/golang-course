package main

import "fmt"

/*
*

	● Using a COMPOSITE LITERAL: ● create a SLICE of TYPE int
	● assign these 10 VALUES 42,43,44,45,46,47,48,49,50,51
	● Range over the slice and print the values out.
	       ○ printout the VALUE and the TYPE
*/
func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} // Slice of type int

	for i, v := range xi {
		fmt.Printf("type: %T - value %v \n", i, v)
	}
}
