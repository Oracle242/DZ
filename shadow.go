package main

import "fmt"

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
	film        [10]cinema
}

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
	for {
		switch vibor() {
		case 1:
			cart = cart.fillCartoteka()
		case 2:
			cart.osnov()
		case 3:
			cart.filmsAmount = 0
			cart.film = [10]cinema{}

		case 4:
			return
		}
	}
}

func (cart cartoteka) fillCartoteka() cartoteka {
	newFilm := cinema{}
	newFilm = newFilm.fillCinema()
	cart.filmsAmount++
	switch cart.filmsAmount {
	case 1:
		cart.film[0] = newFilm
	case 2:
		cart.film[1] = newFilm
	case 3:
		cart.film[2] = newFilm
	case 4:
		cart.film[3] = newFilm
	case 5:
		cart.film[4] = newFilm
	case 6:
		cart.film[5] = newFilm
	case 7:
		cart.film[6] = newFilm
	case 8:
		cart.film[7] = newFilm
	case 9:
		cart.film[8] = newFilm
	case 10:
		cart.film[9] = newFilm
	}
	return cart
}

func (cart cartoteka) osnov() cartoteka {
	fmt.Println("Всего занесено фильмов:", cart.filmsAmount)
	fmt.Println("Ваши фильмы:", cart.film[0].name, cart.film[0].director.name,
		cart.film[1].name, cart.film[1].director.name,
		cart.film[2].name, cart.film[2].director.name,
		cart.film[3].name, cart.film[3].director.name,
		cart.film[4].name, cart.film[4].director.name,
		cart.film[5].name, cart.film[5].director.name,
		cart.film[6].name, cart.film[6].director.name,
		cart.film[7].name, cart.film[7].director.name,
		cart.film[8].name, cart.film[8].director.name,
		cart.film[9].name, cart.film[9].director.name)
	fmt.Println("Выбор фильма от 1 до 10. 11 вернуться в меню.")
	doubleMenu := 0
	fmt.Scanln(&doubleMenu)
	switch doubleMenu {
	case 1:
		fmt.Println(cart.film[0])
	case 2:
		fmt.Println(cart.film[1])
	case 3:
		fmt.Println(cart.film[2])
	case 4:
		fmt.Println(cart.film[3])
	case 5:
		fmt.Println(cart.film[4])
	case 6:
		fmt.Println(cart.film[5])
	case 7:
		fmt.Println(cart.film[6])
	case 8:
		fmt.Println(cart.film[7])
	case 9:
		fmt.Println(cart.film[8])
	case 10:
		fmt.Println(cart.film[9])
	case 11:
		break
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
