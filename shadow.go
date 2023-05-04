package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cinema struct {
	name     string
	age      int
	director portfolio
}

type portfolio struct {
	name     string
	lastName string
	yearsOld int
}

type cartoteka struct {
	filmsAmount int
	films       []cinema
}

// печать меню выбора
func vibor() int {
	numberComand := 0
	fmt.Println("1-Добавить новый фильм в картотеку")
	fmt.Println("2-Посмотреть информацию о фильме")
	fmt.Println("3-Очистить картотеку")
	fmt.Println("4-Выход")
	fmt.Scanln(&numberComand)
	return numberComand
}

func main() {
	cart := cartoteka{}
	cart = cart.openAndReadFile()
	for {
		switch vibor() {
		case 1:
			cart = cart.fillCartoteka()
		case 2:
			cart.osnov()
		case 3:
			cart.filmsAmount = 0
			cart.films = []cinema{}
		case 4:
			cart.save()
			return
		}
	}
}

// Просмотр занесенных фильмов общее меню и в разрезе по фильму
func (cart cartoteka) osnov() cartoteka {
	fmt.Println("Всего занесено фильмов:", cart.filmsAmount)
	fmt.Println("Ваши фильмы:")
	for i := 0; i < len(cart.films); i++ {
		fmt.Println(i+1, ".", cart.films[i].name, cart.films[i].director.name)
	}
	fmt.Println("Выбор фильма от 1 до 10. 11 вернуться в меню.")
	doubleMenu := 0
	fmt.Scanln(&doubleMenu)
	if doubleMenu <= 10 {
		fmt.Println(cart.films[doubleMenu-1])
	}
	return cart
}

// чтение данных в программу
func (cart cartoteka) openAndReadFile() cartoteka {
	file, err := os.Open("Kartoteka.txt")
	if err != nil {
		createFile()
	}
	newFilm := cinema{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line) // Разбиение строки на слова
		for range words {
			newFilm.name = words[0]

			newFilm.director.name = words[2]

			newFilm.director.lastName = words[3]

			num, err := strconv.Atoi(words[1])
			if err != nil {
				log.Fatal(err)
			}
			newFilm.age = num

			num1, error := strconv.Atoi(words[4])
			if error != nil {
				log.Fatal(err)
			}
			newFilm.director.yearsOld = num1
		}
		defer file.Close()
		cart.films = append(cart.films, newFilm)
		cart.filmsAmount++
	}
	return cart
}

// Добавление нового фильма
func (cart cartoteka) fillCartoteka() cartoteka {
	newFilm := cinema{}
	newFilm = newFilm.fillCinema()
	cart.films = append(cart.films, newFilm)
	cart.filmsAmount++
	return cart
}

// Чтение данных от пользователя
func (newFilm cinema) fillCinema() cinema {
	fmt.Println("Введите название фильм")
	fmt.Scanln(&newFilm.name)

	fmt.Println("Введите год фильма")
	fmt.Scanln(&newFilm.age)

	fmt.Println("Введите имя автора фильма")
	fmt.Scanln(&newFilm.director.name)

	fmt.Println("Введите фамилию автора ")
	fmt.Scanln(&newFilm.director.lastName)

	fmt.Println("Введите возраст автора ")
	fmt.Scanln(&newFilm.director.yearsOld)
	return newFilm
}

// Создание файла
func createFile() (*os.File, error) {
	file, err := os.Create("Kartoteka.txt")
	if err != nil {
		log.Fatalf("Ошибка создания файла %s", err)
	}
	return file, err
}

// Запись и сохранение данных в файл
func (cart cartoteka) save() {
	file, err := createFile()
	if err != nil {
		log.Fatalf("Ошибка создпния файла %s", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range cart.films {
		age := strconv.Itoa(line.age)
		yearsOld := strconv.Itoa(line.director.yearsOld)
		_, err = writer.WriteString(line.name + " " + age + " " + line.director.name + " " + line.director.lastName + " " + yearsOld + "\n")
		if err != nil {
			log.Fatalf("ошибка записи данных %s", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Ошибка записи в буфер %s", err)
	}
}
