package main

import "fmt"

func main() {
	xs := []string{
		"AlmondBiscottiCafé",
		"BalsamicStrawberry(GF)",
		"Bittersweet Chocolate(GF)",
	}

	fmt.Println(xs)
	fmt.Println(len(xs))

	for _, v := range xs {
		fmt.Printf("%#v\n", v)
	}

	fmt.Println("-------------------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
}
