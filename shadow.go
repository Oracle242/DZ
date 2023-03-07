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
	cart.film[cart.filmsAmount] = newFilm
	cart.filmsAmount++
	return cart
}

func (cart cartoteka) osnov() cartoteka {
	fmt.Println("Всего занесено фильмов:", cart.filmsAmount)
	fmt.Println("Ваши фильмы:")
	for i := 0; i < 9; i++ {
		fmt.Println(i, ".", cart.film[i].name, cart.film[i].director.name)
	}
	fmt.Println("Выбор фильма от 1 до 10. 11 вернуться в меню.")
	doubleMenu := 0
	fmt.Scanln(&doubleMenu)
	switch doubleMenu < 10 {
	case true:
		fmt.Println(cart.film[doubleMenu-1])
	case false:
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
