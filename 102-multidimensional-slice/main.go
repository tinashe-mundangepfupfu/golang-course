package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	jm := []string{"Jenny", "Moneypenny", "Wolverine Tracks", "Guiness"}

	fmt.Println(jb)
	fmt.Println(jm)

	xxs := [][]string{jb, jm} //slice storing slices of string

	fmt.Println(xxs)
}
