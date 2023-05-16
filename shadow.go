package main

import (
	"fmt"
	"myProj/electronic"
)

func printCharacteristics(p electronic.Phone) {
	fmt.Printf("Brand: %s\nModel: %s\nType: %s\n", p.Brand(), p.Model(), p.Type())
	if sp, ok := p.(electronic.Smartphone); ok {
		fmt.Printf("OS: %s\n", sp.OS())
	}
	if stp, ok := p.(electronic.StationPhone); ok {
		fmt.Printf("Buttons: %d\n", stp.ButtonsCount())
	}
}

func main() {
	apple := electronic.NewApplePhone("iPhone X")
	android := electronic.NewAndroidPhone("Samsung", "Galaxy S9")
	radio := electronic.NewRadioPhone("Sony", 12)

	printCharacteristics(apple)
	fmt.Println()
	printCharacteristics(android)
	fmt.Println()
	printCharacteristics(radio)
}
