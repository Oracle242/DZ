package main

import "fmt"

// сортировка на символы
func sorting(sumols string) {
	for i := 0; i < len(sumols)-1; i++ {
		if sumols[i] != sumols[i+1] {
			archive(sumols, string(sumols[i]))
		}
	}
	archive(sumols, string(sumols[len(sumols)-1]))
	fmt.Printf("\n")
}

// подсчет повторений символов
func archive(sumols, sumol string) {
	r := 0
	for i := 0; i < len(sumols); i++ {
		if sumol == string(sumols[i]) {
			r++
		}
	}
	fmt.Printf("%v%v", sumol, r)
}

func main() {
	sumols := "aaabbbbccccccr"
	fmt.Println("Укажите строку для проверки на архива")
	sorting(sumols)
}
