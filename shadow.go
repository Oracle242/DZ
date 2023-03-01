package main

import (
	"fmt"
	"strings"
)

/*
1. fio := "Vasya Pupkin" получить имя.
var fios [5]string = [5]string{"Vasya Pupkin", "Vasya Pupkin", "Vasya Pupkin", "Vasya Pupkin", "Vasya Pupkin"}
отсортировать имена двумя способами - по возрастанию и уыванию
*/
func main() {
	var fios [5]string = [5]string{"Vasya Pupkin", "Sasha Dupkin", "Pahy Sumkin", "Dasha Lupkin", "Natasha Turkin"}
	var name [5]string
	for i := 0; i < 5; i++ {
		arr := strings.Split(fios[i], " ")
		name[i] = arr[0]
	}

	for j := 0; j < 4; j++ {
		for a := 0; a < 4; a++ {
			if name[a] > name[a+1] {
				nameShadow := name[a]
				name[a] = name[a+1]
				name[a+1] = nameShadow
			}
		}

	}
	fmt.Println(name)
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
