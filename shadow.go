package main

import (
	"fmt"
)

func main() {
	var numb [5]int = [5]int{32, -87, 50, 1, 6}

	for i := 0; i < 5; i++ {
		fmt.Scanln(&numb[i])
	}
	fmt.Println(numb)
	for j := 0; j < 5; j++ {
		for i := 0; i < 4; i++ {
			if numb[i] > numb[i+1] {
				numbShadow := numb[i]
				numb[i] = numb[i+1]
				numb[i+1] = numbShadow
			}
		}
	}
	fmt.Println(numb)
}
