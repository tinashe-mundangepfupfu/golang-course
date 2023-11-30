package main

import "fmt"

var myTestVar = "my package var"

const (
	testConst = "TEST_CONST"
)

func main() {
	insideMain := "we are in main"

	fmt.Println(myTestVar)
	fmt.Println(testConst)
	fmt.Println(insideMain)

}
