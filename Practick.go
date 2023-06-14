package main

import "fmt"

func searchPrimeNumbers(numb int) bool {
	for i := numb - 1; i > 1; i-- {
		if numb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	counter := 0
	for i := 2; counter < 20; i++ {
		elNu := searchPrimeNumbers(i)
		if elNu == true {
			fmt.Println(i)
			counter++
		}
	}
}
