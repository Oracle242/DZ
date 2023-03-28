package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("Фильмотека.txt")
	if err != nil {
		fmt.Printf("Возникла ошибка создания библиотеки фильмов [%s]", err)
		return
	}
	fmt.Printf("Файл с названием %s создан", file.Name())
}

// Изучить раздел misc.
// Практика слайсов.
