package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func PeopleMap() map[string]Man {
	people := make(map[string]Man)
	people["Алексей"] = Man{Name: "Алексей", LastName: "Руманов", Age: 22, Gender: "Мужчина", Crimes: 1}
	people["Петр"] = Man{Name: "Петр", LastName: "Пермитин", Age: 65, Gender: "Мужчина", Crimes: 3}
	people["Мария"] = Man{Name: "Мария", LastName: "Краснова", Age: 23, Gender: "Женщина", Crimes: 5}
	people["Роман"] = Man{Name: "Роман", LastName: "Гудман", Age: 37, Gender: "Мужчина", Crimes: 2}
	people["Софья"] = Man{Name: "Софья", LastName: "Гудман", Age: 34, Gender: "Женщина", Crimes: 11}
	return people
}

func searchPeopleCrimes() {
	nameCr := []string{"Петр", "Алексей", "Роман"}
	var person Man
	people := PeopleMap()
	mainZlodey := people[nameCr[0]]
	for i := 0; i < len(nameCr); i++ {
		person = people[nameCr[i]]

		if person.Crimes > mainZlodey.Crimes {
			mainZlodey = person
		}
	}
	printBd(mainZlodey)
}

func printBd(mainZlodey Man) {
	if mainZlodey.Crimes == 0 {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	} else {
		fmt.Printf("Самый главный злодей:\n Имя: %s\n Фамилия: %s\n Возраст: %v\n Пол: %s\n Колличество преступлений: %v\n", mainZlodey.Name, mainZlodey.LastName, mainZlodey.Age, mainZlodey.Gender, mainZlodey.Crimes)
	}
}

func main() {
	searchPeopleCrimes()
}
