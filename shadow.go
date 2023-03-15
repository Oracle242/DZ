package main

import "fmt"

func main() {
	nameArr := [6]string{"vasya", "petya", "masha", "vasya", "vasya", "fedya"}
	slov := map[string]int{}

	// var mass []string
	// var mass1 []int

	// for key, valye := range slov {
	// 	mass = append(mass, key)
	// 	mass1 = append(mass1, valye)
	// }

	fmt.Println(nameArr)
	// fmt.Println(mass)
	// fmt.Println(mass1)
}

// 1.
// ['vasya', 'petya', 'masha', 'vasya', 'vasya', 'fedya']
// посчитать с помощью map сколько раз встречается каждое имя
//

// 2.
// ['vasya', 'petya', 'masha', 'vasya', 'vasya', 'fedya']
// вывести уникальные имена
// отсортировать имена
// разобраться с range(с разными вариациями)
// Изучить раздел misc.
