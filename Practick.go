package main

import "fmt"

func findMostOftenRepeatedOptimized(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 1
		for j := i + 1; j < len(array); j++ {
			if number == array[j] {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}

	return array[maxIndex], nil
}

func main() {
	num := []int{1, 7, 2, 3, 6, 4, 3, 2, 3, 1, 3}
	fmt.Println(findMostOftenRepeatedOptimized(num))
}
