package main

import (
	"fmt"
	"myProj/autosolon"
)

func printCharacteristics(a autosolon.Auto) {
	fmt.Printf("Brand: %s\nModel: %s\nDimensions: %v\nMaxSpeed: %v\nEnginePower: %v\n\n", a.Brand(), a.Model(), a.Dimensions(), a.MaxSpeed(), a.EnginePower())
}

func main() {
	bmw := autosolon.NewBMWEuropean("X5", autosolon.NewCMDimensions(490.7, 200.6, 175.3), 210, 300)
	mercedes := autosolon.NewMercedesEuropean("S-Class", autosolon.NewCMDimensions(490.7, 200.6, 175.3), 210, 300)
	dodge := autosolon.NewDodgeEuropean("Charger", autosolon.NewInchesDimensions(198.4, 75.0, 58.2), 120, 292)

	printCharacteristics(bmw)
	printCharacteristics(mercedes)
	printCharacteristics(dodge)
}
