package main

import "fmt"

func searchPrimeNumbers(numb int) int {
	for i := numb - 1; i > 1; i-- {
		if numb%i == 0 {
			return 0
		}
	}
	return numb
}

func main() {
	counter := 0
	for i := 2; counter < 20; i++ {
		elNu := searchPrimeNumbers(i)
		if elNu != 0 {
			fmt.Println(elNu)
			counter++
		}
	}
}
