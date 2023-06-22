package main

import "fmt"

func trimLessAverage(array []int) ([]int, error) {
	array2 := []int{}
	total, average := 0, 0

	for _, numb := range array {
		total += numb
	}

	if total > len(array) {
		average = total / len(array)
	} else {
		return array2, fmt.Errorf("Среднее число не высчитать")
	}

	for i := 0; i < len(array); i++ {
		if array[i] > average {
			array2 = append(array2, array[i])
		}
	}
	return array2, nil
}

func main() {
	num := []int{1, 2}
	fmt.Println(trimLessAverage(num))
}
