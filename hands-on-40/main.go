package main

import "fmt"

func main() {
	afood := [...]string{
		"AlmondBiscottiCafé",
		"BalsamicStrawberry(GF)",
		"Bittersweet Chocolate(GF)",
		"BlueberryPancake(GF)",
		"BourbonTurtle(GF)",
		"BrownedButter CookieDough",
		"BananaPudding",
		"ChocolateCoveredBlackCherry(GF)",
		"ChocolateFudgeBrownie",
		"ChocolatePeanutButter(GF)",
		"Coffee(GF)",
		"Cookies&Cream",
		"FreshBasil (GF)",
		"GardenMintChip(GF)",
		"LavenderLemonHoney(GF)",
		"LemonBar",
		"Madagascar Vanilla(GF)",
		"Matcha(GF)",
		"MidnightChocolateSorbet(GF,V)",
		"Non-DairyChocolate PeanutButter(GF,V)",
		"Non-DairyCoconutMatcha(GF,V)", "Non-DairySunbutter Cinnamon(GF,V)",
		"OrangeCream(GF)", "PeanutButterCookieDough", "Raspberry Sorbet(GF,V)",
		"SaltyCaramel (GF)", "SlateSlateDifferent",
		"StrawberryLemonade Sorbet(GF,V)", "VanillaCaramelBlondie", "VietnameseCinnamon(GF)", "Wolverine Tracks(GF)",
	}
	fmt.Printf("Length of the array is %#v\n ", len(afood))
	fmt.Println(afood)

}
