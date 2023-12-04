package main

import "fmt"

/*
*
 */
func main() {

	xStates := make([]string, 0, 50)

	xStates = append(xStates, `Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, ` Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, ` Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, ` Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `NewHampshire`, `NewJersey`, `NewMexico`, `NewYork`, `NorthCarolina`, `NorthDakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `RhodeIsland`, `SouthCarolina`, `SouthDakota`, `Tennessee`, `Texas`, ` Utah`, `Vermont`, `Virginia`, `Washington`, `WestVirginia`, `Wisconsin`, `Wyoming`)

	fmt.Print(len(xStates), cap(xStates))
	fmt.Println("-----------------------")

	for i := 0; i < cap(xStates); i++ {
		fmt.Printf("index: %v - value: %v \n ", i, xStates[i])
	}

}
