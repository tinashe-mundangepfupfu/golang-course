package main

import "fmt"

type Person struct {
	first string
}

type SecretAgent struct {
	Person
	ltk bool
}

func (p Person) speak() { // how you create a method
	fmt.Println("I am", p.first)
}
func (sa SecretAgent) speak() { // how you create a method
	fmt.Println("I am secret agent", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := SecretAgent{
		Person: Person{
			first: "James",
		},
		ltk: true,
	}

	p2 := Person{first: "Jenny"}

	//sa1.speak()
	//p2.speak()

	saySomething(sa1)
	saySomething(p2)

}
