package main

import "fmt"

func main() {
	nameArr := [10]string{"vasya", "petya", "fedya", "vasya", "vasya", "fedya", "masha", "petya", "petya", "alex"}
	slais := []string{}

	for i := 0; i < len(nameArr); i++ {
		max := nameArr[0]
		min2 := 0
		for j := 0; j < len(nameArr); j++ {
			if nameArr[j] < max {
				max = nameArr[j]
				min2 = j
			}
		}
		nameArr[min2] = "zzz"
		slais = append(slais, max)
	}

	fmt.Println(slais)
}

/*
44 отсортировать массив с помощью метода выбором (где мы ищем максимальные элементы)
 и отсортированные значения помещать в слайс
*/

// разобраться с range(с разными вариациями)
// Изучить раздел misc.
// Практика слайсов.
