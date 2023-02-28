package main

import (
	"fmt"
)

func main() {
	var numb [5]int = [5]int{27, 1, 12, 2, 260}
	var numb2 [5]int = [5]int{0, 0, 0, 0, 0}

	fmt.Println(numb)
	fmt.Println(numb2)

	for j := 0; j < 5; j++ {
		for i := 0; i < 4; i++ {
			if numb[i] > numb[i+1] {
				numbShadow := numb[i]
				numb[i] = numb[i+1]
				numb[i+1] = numbShadow
			}
		}
	}

	for f := 0; f < 5; f++ {
		if numb[f] > numb2[f] {
			numbShadow := numb[f]
			numb[f] = numb2[f]
			numb2[f] = numbShadow
			numb[f] = 999
		}
	}

	fmt.Println(numb)
	fmt.Println(numb2)
}

/*
1. fio := "Vasya Pupkin" получить имя.
var fios [5]string = [5]string{"Vasya Pupkin", "Vasya Pupkin", "Vasya Pupkin", "Vasya Pupkin", "Vasya Pupkin"}
отсортировать имена двумя способами - по возрастанию и уыванию
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
