package main

import "fmt"

func palindromeSearch(num string) bool {
	if num < "1" {
		return false
	}
	for i := 1; i < len(num)+1; i++ {
		if num[i-1] != num[len(num)-i] {
			return false
		}
	}
	return true
}

func main() {
	numb := "765"
	fmt.Println("Укажите число для проверки на палиндром")
	fmt.Scanln(&numb)
	fmt.Println(palindromeSearch(numb))
}
