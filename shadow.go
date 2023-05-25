package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func (m Man) NewName() string {
	return m.Name
}
func (m Man) NewLastName() string {
	return m.LastName
}
func (m Man) NewAge() int {
	return m.Age
}
func (m Man) NewGender() string {
	return m.Gender
}
func (m Man) NewCrimes() int {
	return m.Crimes
}

func NewPeopleMap() map[string]Man {
	people := make(map[string]Man)
	people["Алексей"] = Man{Name: "Алексей", LastName: "Руманов", Age: 22, Gender: "Мужчина", Crimes: 1}
	people["Петр"] = Man{Name: "Петр", LastName: "Пермитин", Age: 65, Gender: "Мужчина", Crimes: 2}
	people["Мария"] = Man{Name: "Мария", LastName: "Краснова", Age: 23, Gender: "Женщина", Crimes: 5}
	people["Роман"] = Man{Name: "Роман", LastName: "Гудман", Age: 37, Gender: "Мужчина", Crimes: 6}
	people["Софья"] = Man{Name: "Софья", LastName: "Гудман", Age: 34, Gender: "Женщина", Crimes: 11}
	return people
}

func Crimess() {
	key := []string{}
	crimes := []int{}
	for k, v := range NewPeopleMap() {
		if v.Crimes >= 0 {
			key = append(key, k)
			crimes = append(crimes, v.Crimes)
		}
	}
	max, name := findMinAndMax(crimes, key)
	if max == 0 {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	} else {
		for k, v := range NewPeopleMap() {
			if k == name {
				fmt.Printf("Самый главный злодей:\n Имя: %s\n Фамилия: %s\n Возраст: %v\n Пол: %s\n Колличество преступлений: %v\n", v.Name, v.LastName, v.Age, v.Gender, v.Crimes)
			}
		}
	}
}

func findMinAndMax(crimes []int, key []string) (int, string) {
	var max int
	var name string
	if len(crimes) == 1 {
		return crimes[0], key[0]
	}
	if crimes[0] > crimes[1] {
		max = crimes[0]
		name = key[0]
	} else {
		max = crimes[1]
		name = key[1]
	}
	for i := 2; i < len(crimes); i++ {
		if crimes[i] > max {
			max = crimes[i]
			name = key[i]
		}
	}
	return max, name
}

func main() {
	Crimess()
}
