package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	film := "Терминатор"
	file, err := os.Create("Фильмотека.txt")
	if err != nil {
		fmt.Printf("Возникла ошибка создания библиотеки фильмов [%s]", err)
		return
	}

	stringlen, err := file.WriteString(film)
	if err != nil {
		fmt.Printf("Возникла ошибка записи в библиотеку фильмов [%s]", err)
		return
	}

	openFile, err := os.Open(file.Name())
	if err != nil {
		fmt.Printf("Возникла ошибка прочтения библиотеки фильмов [%s]", err)
		return
	}

	data := make([]byte, stringlen)

	for {
		_, err := openFile.Read(data)
		if err != io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Ошибка чтения из файла [%s]", err)
			break
		}
	}
	fmt.Printf(string(data))
}

// Изучить раздел misc.
// Практика слайсов.
