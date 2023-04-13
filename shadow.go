package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	film        []cinema
}

func vibor() int {
	numberComand := 0
	fmt.Println("1-Добавить новый фильм в картотеку")
	fmt.Println("2-Посмотреть информацию о фильме")
	fmt.Println("3-Очистить картотеку")
	fmt.Println("4-Запись в файл txt")
	fmt.Println("5-Выход")
	fmt.Scanln(&numberComand)
	return numberComand
}

func main() {
	cart := cartoteka{}
	for {
		switch vibor() {
		case 1:
			cart = cart.fillCartoteka()
		case 2:
			cart.osnov()
		case 3:
			cart.filmsAmount = 0
			cart.film = []cinema{}
		case 4:
			cart.save()
		case 5:
			return
		}
	}
}

func (cart cartoteka) fillCartoteka() cartoteka {
	newFilm := cinema{}
	newFilm = newFilm.fillCinema()
	cart.film = append(cart.film, newFilm)
	cart.filmsAmount++
	return cart
}

func (cart cartoteka) osnov() cartoteka {
	fmt.Println("Всего занесено фильмов:", cart.filmsAmount)
	fmt.Println("Ваши фильмы:")
	for i := 0; i < len(cart.film); i++ {
		fmt.Println(i, ".", cart.film[i].name, cart.film[i].director.name)
	}
	fmt.Println("Выбор фильма от 1 до 10. 11 вернуться в меню.")
	doubleMenu := 0
	fmt.Scanln(&doubleMenu)
	if doubleMenu <= 10 {
		fmt.Println(cart.film[doubleMenu-1])
	}
	return cart
}

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

func (cart cartoteka) save() {
	bufer := []cinema{}

	createFile, err := os.Create("Картотека.txt")
	if err != nil {
		log.Fatalf("Ошибка открытия файла %s", err)
	}
	defer createFile.Close()
	writer := bufio.NewWriter(createFile)

	buf := cart.dataProcessing(bufer)
	for _, line := range buf {
		age1 := line.age
		age2 := strconv.Itoa(age1)
		yearsOld1 := line.director.yearsOld
		yearsOld2 := strconv.Itoa(yearsOld1)
		_, err = writer.WriteString(line.name + " " + age2 + " " + line.director.name + " " + line.director.lastName + " " + yearsOld2 + "\n")
		if err != nil {
			log.Fatalf("ошибка считывания данных %s", err)
		}

		err = writer.Flush()
		if err != nil {
			fmt.Printf("Ошибка записи в буфер %s", err)
		}
	}
}
func (cart cartoteka) dataProcessing(bufer []cinema) []cinema {

	for i := range cart.film {
		bufer = append(bufer, cart.film[i])
	}
	return bufer
}
