package main

import "fmt"

func main() {
	var number [5]int
	for i := 0; i < 5; i++ {
		fmt.Scanln(&number[i])
	}

	min := number[0]
	imin := 0
	max := number[0]
	imax := 0
	for i := 0; i < 5; i++ {
		if number[i] < min {
			min = number[i]
			imin = i
		} else if number[i] > max {
			max = number[i]
			imax = i
		}
	}

	changing := number[imin]
	number[imin] = number[imax]
	number[imax] = changing

	fmt.Println(number)
}
