package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fale := []string{}
	znak := " !"
	strSlice := []string{}
	intSlice := []int{}

	pathFile := "C:/Users/PSXnv/Desktop/Task.txt"

	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fale = append(fale, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Ошибка при чтении файла: %s", err)
	}
	file.Close()

	for i := 0; i < len(fale); i++ {
		if fale[i] < "0" || fale[i] > "999" {
			strSlice = append(strSlice, fale[i]+znak)
		} else {
			numbr, err := strconv.Atoi(fale[i])
			if err != nil {
				log.Fatal(err)
			}
			intSlice = append(intSlice, numbr*2)
		}
	}

	createFile, err := os.Create("random.txt")
	if err != nil {
		fmt.Printf("Ошибка открытия файла %s", err)
		return
	}
	defer createFile.Close()

	writer := bufio.NewWriter(createFile)

	for _, str := range strSlice {
		_, err = writer.WriteString(str + "\n")
		if err != nil {
			fmt.Printf("ошибка считывания данных %s", err)
			return
		}
	}
	for _, num := range intSlice {
		_, err = writer.WriteString(strconv.Itoa(num) + "\n")
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

// Изучить раздел misc.
// Практика слайсов.
