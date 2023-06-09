package main

import "fmt"

func ellementaruNumb(numbs int) int {
	for i := numbs - 1; i > 1; i-- {
		if numbs%i == 0 {
			return 0
		}
	}

	return numbs
}

func main() {
	x := 0
	for i := 0; x <= 20; i++ {
		ellementaruNumb(i)
		if ellementaruNumb(i) != 0 {
			fmt.Println(ellementaruNumb(i))
			x++
		}
	}
}
