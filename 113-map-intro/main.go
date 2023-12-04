package main

import "fmt"

func main() {

	// Creating a Map
	// ----- 1
	am := map[string]int{
		"Todd":  42,
		"Henry": 16,
		"Paget": 14,
	}

	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v \n", am)

	// ----- 2
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 27

	fmt.Println(an)
	fmt.Printf("%#v \n", an)

	fmt.Println(len(an))

}
