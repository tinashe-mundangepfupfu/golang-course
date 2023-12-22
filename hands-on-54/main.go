package main

import "fmt"

/*
 *
 */

type person struct {
	firstName         string
	lastName          string
	xIceCreamFlavours []string
}

func main() {
	p1 := person{
		firstName: "Paul",
		lastName:  "Atreidis",
		xIceCreamFlavours: []string{
			"Vanilla",
			"Blueberry",
			"chocolate",
		},
	}

	p2 := person{
		firstName: "Tom",
		lastName:  "Sawyer",
		xIceCreamFlavours: []string{
			"Cherry",
			"Banana",
			"Lemon",
		},
	}

	mp := make(map[string]person)

	mp[p1.lastName] = p1

	mp[p2.lastName] = p2

	for k, v := range mp {
		fmt.Println(k, v)
	}

	//fmt.Println("p1 is: ", p1.firstName, p1.lastName)
	//fmt.Println("Fav flavour:")
	//for _, flavour := range p1.xIceCreamFlavours {
	//	fmt.Println(flavour)
	//}
	//
	//fmt.Println("-----------------------")
	//
	//fmt.Println("p2 is: ", p2.firstName, p2.lastName)
	//fmt.Println("Fav flavour:")
	//for _, flavour := range p2.xIceCreamFlavours {
	//	fmt.Println(flavour)
	//}

}
