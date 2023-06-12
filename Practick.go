package main

import "fmt"

func ellementaruNumb(numb int) int {
	for i := numb - 1; i > 1; i-- {
		if numb%i == 0 {
			return 0
		}
	}
	return numb
}

func main() {
	x := 0
	for i := 0; x < 3; i++ {
		elNu := ellementaruNumb(i)
		if elNu != 0 {
			fmt.Println(elNu)
			x++
		}
	}
}
