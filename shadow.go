package main

import (
	"fmt"
)

func main() {
	var namb [5]int = [5]int{27, 88, 12, 54, 1}
	var nambCore [5]int = [5]int{0, 0, 0, 0, 0}

	min := namb[0]
	minid := 0

	for id, element := range namb {
		if element < min {
			nambCore[minid] = element
			min = id
		}
	}
	namb[min] = 999

	fmt.Println(min)
	fmt.Println(namb)
	fmt.Println(nambCore)
}

/*
2. сортировка вставками
имеется массив [27, 1, 12, 2, 260]
создаем новый пустой массив [0, 0, 0, 0, 0]
ищем минимальное, вставляем его на 0 индекс, и в первоначальном массиве заменяем это число на супер большое:
[27, 99999, 12, 2, 260]
[1, 0, 0, 0, 0]
снова ищем минимальное. вставляем его на 1 место и заменяем
[27, 99999, 12, 99999, 260]
[1, 2, 0, 0, 0]
*/
