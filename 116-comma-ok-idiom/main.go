package main

import "fmt"

func main() {
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

	delete(an, "Steph")

	// For range over a map
	for k, v := range an {
		fmt.Println(k, v)
	}

	for _, v := range an {
		fmt.Println(v)
	}

	for k := range an {
		fmt.Println(k)
	}

}
