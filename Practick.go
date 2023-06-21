package main

import "fmt"

func findMostOftenRepeatedWithMap(array []int) (max int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("не удалось найти повторяющиеся числа в пустом срезе")
	}
	mapps := make(map[int]int)
	max = 0
	repetition := 0
	for i := 0; i < len(array); i++ {
		mapps[array[i]] += 1
		if mapps[array[i]] > repetition {
			repetition = mapps[array[i]]
		}
		if mapps[array[i]] == repetition {
			max = array[i]

		}
	}
	return max, fmt.Errorf("Ошибки не обнаружены!")
}

func main() {
	num := []int{7, 7, 7, 1, 8, 2, 6, 6, 6, 8, 3, 8, 2, 8, 1, 9, 9, 9}
	fmt.Println(findMostOftenRepeatedWithMap(num))
}
