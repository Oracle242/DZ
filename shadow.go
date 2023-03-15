package main

import "fmt"

func main() {
	slov := make(map[string]string)
	nameArr := [10]string{"vasya", "petya", "fedya", "vasya", "vasya", "fedya", "masha", "petya", "petya", "vasya"}
	for _, v := range nameArr {
		slov[v] = ""
	}
	fmt.Println(slov)
}

/*
1.
['vasya', 'petya', 'masha', 'vasya', 'vasya', 'fedya']
посчитать с помощью map сколько раз встречается каждое имя

Печать из терминала - map[fedya:2 masha:1 petya:3 vasya:4]
*/
/*
2.
['vasya', 'petya', 'masha', 'vasya', 'vasya', 'fedya']
вывести уникальные имена
отсортировать имена
Печать из терминала - map[fedya: masha: petya: vasya:]
*/

// разобраться с range(с разными вариациями)
// Изучить раздел misc.
