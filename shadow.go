package main

import "fmt"

func main() {
	nameArr := [10]string{"vasya", "petya", "fedya", "vasya", "vasya", "fedya", "masha", "petya", "petya", "alex"}
	slais := []string{}
	for i := 0; i < len(nameArr); i++ {
		for j := 0; j < len(nameArr)-1; j++ {
			if nameArr[j] > nameArr[j+1] {
				nameShadow := nameArr[j]
				nameArr[j] = nameArr[j+1]
				nameArr[j+1] = nameShadow
			}
		}
	}

	for t := 0; t < len(nameArr); t++ {
		slais = append(slais, nameArr[t])
	}

	fmt.Println(slais)
}

/*
1.
 отсортировать массив ["vasya", "petya", "fedya", "vasya", "vasya", "fedya", "masha", "petya", "petya", "alex"]
 с помощью метода выбором (где мы ищем максимальные элементы)
 и отсортированные значения помещать в слайс
*/

// разобраться с range(с разными вариациями)
// Изучить раздел misc.
// Практика слайсов.
