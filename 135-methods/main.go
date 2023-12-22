package main

import "fmt"

type Person struct {
	first string
}

func (p Person) speak() { // how you create a method
	fmt.Println("I am", p.first)
}

func main() {
	p1 := Person{first: "James"}
	p2 := Person{first: "Jenny"}

	p1.speak()
	p2.speak()

}
