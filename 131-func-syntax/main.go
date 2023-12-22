package main

import "fmt"

func main() {
	foo()
	bar("Tinashe")
	s := aloha("Tom")
	fmt.Println(s)

	age, dy := dogYears("Tinashe", 50)
	fmt.Println(age, dy)
}

func foo() {
	fmt.Println("I am from foo!")
}

func bar(s string) {
	fmt.Println("My name is ", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years"), age
}
