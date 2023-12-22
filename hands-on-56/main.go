package main

import "fmt"

/*
 *
 */

type Engine struct {
	electric bool
}

type Vehicle struct {
	engine Engine
	make   string
	model  string
	doors  string
	color  string
}

func main() {
	v1 := Vehicle{
		engine: Engine{
			electric: true,
		},
		make:  "BMW",
		model: "520",
		doors: "4",
		color: "white",
	}

	v2 := Vehicle{
		engine: Engine{
			electric: false,
		},
		make:  "Mercedes Benz",
		model: "C200",
		doors: "4",
		color: "Black",
	}

	fmt.Println("1st car: ", v1)
	fmt.Println("2nd car: ", v2)
	fmt.Println("1st car make", v1.make)
	fmt.Println("1st car make", v1.engine.electric)
	fmt.Println("1st car make", v2.make)

	//mp := make(map[string]person)
	//
	//mp[p1.lastName] = p1
	//
	//mp[p2.lastName] = p2
	//
	//for k, v := range mp {
	//	fmt.Println(k, v)
	//}

}
