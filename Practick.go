package main

import "fmt"

func trimNegative(array []int) (max []int) {

	for i := 0; i < len(array)-1; i++ {
		if array[i] < 0 {
			array = append(array[:i], array[i+1:]...)
		}

	}

	return array
}

func main() {
	num := []int{8, 1, -3, 4}
	fmt.Println(trimNegative(num))
}
