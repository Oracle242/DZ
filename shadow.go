package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	bufer := []string{}
	pathFile := "C:/Users/PSXnv/Desktop/Task.txt"

	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %s", err)
		return
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		bufer = append(bufer, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Ошибка при чтении файла: %s", err)
	}
	file.Close()

	createFile, err := os.Create("random.txt")
	if err != nil {
		fmt.Printf("Ошибка открытия файла %s", err)
		return
	}
	defer createFile.Close()
	writer := bufio.NewWriter(createFile)

	for _, line := range dataProcessing(bufer) {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("ошибка считывания данных %s", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Ошибка записи в буфер %s", err)
	}
}

func dataProcessing(bufer []string) []string {
	znak := " !"
	for i, line := range bufer {
		if line < "0" || line > "999" {
			bufer[i] = line + znak
		} else {
			numbr, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			bufer[i] = strconv.Itoa(numbr * 5)
		}
	}
	return bufer
}

// Изучить раздел misc.
// Практика слайсов.
