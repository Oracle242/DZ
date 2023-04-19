package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{77, 32, -5, 677, -99, 223, -34, 78, 92, 13}
	bubbleSort(arr)
	fmt.Println("Sorted array:", arr)
}
