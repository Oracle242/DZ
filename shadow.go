package main

import "fmt"

<<<<<<< HEAD
type Numbers struct {
	num1 int
	num2 int
}

type NumbersInterface interface {
	Sum() int
	Multiplay() int
	Division() float64
	Substract() int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}
func (n Numbers) Substract() int {
	return n.num1 - n.num2
}
func (n Numbers) Multiplay() int {
	return n.num1 * n.num2
}
func (n Numbers) Division() float64 {
	if n.num2 == 0 {
		fmt.Println("На ноль делить низя")
	}
	return float64(n.num1) / float64(n.num2)
}
func main() {

	numb1 := 0
	numb2 := 0
	sign := "0"
	fmt.Println("Введи первое число")
	fmt.Scanln(&numb1)
	fmt.Println("Введи +-*/")
	fmt.Scanln(&sign)
	fmt.Println("Введи второе число")
	fmt.Scanln(&numb2)

	var i NumbersInterface
	numbers := Numbers{numb1, numb2}
	i = numbers

	switch sign {
	case "+":
		fmt.Printf("Ответ: %d\n", i.Sum())
	case "-":
		fmt.Printf("Ответ: %d\n", i.Substract())
	case "*":
		fmt.Printf("Ответ: %d\n", i.Multiplay())
	case "/":
		fmt.Printf("Ответ: %f\n", i.Division())
	}
}
=======
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
			cart.film = []cinema{}
		case 4:
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

	// fmt.Println("Введите год фильма")
	// fmt.Scanln(&newFilm.age)

	// fmt.Println("Введите имя автора фильма")
	// fmt.Scanln(&newFilm.director.name)

	// fmt.Println("Введите фамилию автора ")
	// fmt.Scanln(&newFilm.director.lastName)

	// fmt.Println("Введите возраст автора ")
	// fmt.Scanln(&newFilm.director.yearsOld)
	return newFilm
}
>>>>>>> 0bf894e6d8b8e2e1db263da3276beec647c3675f
