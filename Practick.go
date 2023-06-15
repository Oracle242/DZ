package main

import "fmt"

// сортировка на символы
func sorting(sumols string) {
	enSt := 0
	for i := 0; i < len(sumols)-1; i++ {
		if sumols[i] != sumols[i+1] {
			enSt = i + 1
			archive(sumols, string(sumols[i]), enSt)
		}
	}
	archive(sumols[enSt:], string(sumols[len(sumols)-1]), len(sumols[enSt:]))
	fmt.Printf("\n")
}

// подсчет повторений символов
func archive(sumols, sumol string, enSt int) {
	r := 0
	for i := 0; i < enSt; i++ {
		if sumol == string(sumols[i]) {
			r++
		}

	}
	fmt.Printf("%v%v", sumol, r)
}

func main() {
	sumols := "aaabbbbccccccaaaa"
	fmt.Println("Укажите строку для проверки на архива")
	sorting(sumols)
}
