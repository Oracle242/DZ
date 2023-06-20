package main

import "fmt"

func findMostOftenRepeatedWithMap(array []int) (max int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("не удалось найти повторяющиеся числа в пустом срезе")
	}
	mapps := make(map[int]int)
	d, smax := 0, 0
	max = 0
	for i := 0; i < len(array); i++ {
		if mapps[array[i]] == 0 {
			mapps[array[i]] = 1

		} else {
			d++
			mapps[array[i]] = d
			smax = d
		}
		if mapps[array[i]] == smax {
			max = array[i]
		}
	}
	return max, fmt.Errorf("Ошибки не обнаружены!")
}

func main() {
	num := []int{2, 8, 3, 8, 2, 8, 1}
	fmt.Println(findMostOftenRepeatedWithMap(num))
}
