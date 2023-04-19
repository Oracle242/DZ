package main

import "fmt"

func findMinAndMax(arr []int) (int, int) {
	var min, max int
	if len(arr) == 1 {
		return arr[0], arr[0]
	}
	if arr[0] > arr[1] {
		max = arr[0]
		min = arr[1]
	} else {
		max = arr[1]
		min = arr[0]
	}
	for i := 2; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}

func main() {
	arr := []int{77, 32, -5, 677, -99, 223, -34, 78, 92, 13}
	min, max := findMinAndMax(arr)
	fmt.Println("Array:", arr)
	fmt.Printf("Min: %d, Max: %d\n", max, min)
}
