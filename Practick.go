package main

import "fmt"

func trimLessAverage(array []int) []int {
	total := 0
	for _, numb := range array {
		total += numb
	}

	array2 := []int{}
	average := total / len(array)

	for i := 0; i < len(array); i++ {
		if array[i] > average {
			array2 = append(array2, array[i])
		}
	}
	return array2
}

func main() {
	num := []int{1, 9, 1, 0, 9, 5, 0, 1}
	fmt.Println(trimLessAverage(num))
}
