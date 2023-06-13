package main

import "fmt"

func palindromeSearch(nubm string) bool {
	if len(string(nubm)) == 1 || nubm <= string("1") {
		return false
	}

	essence := []string{}
	reflection := []string{}
	for i := 0; i <= len(string(nubm))-1; i++ {
		essence = append(essence, string(nubm[i]))
	}

	for j := 1; j <= len(string(nubm)); j++ {
		reflection = append(reflection, string(nubm[len(string(nubm))-j]))
	}

	for i := 0; i < len(essence)-1; i++ {
		if essence[i] != reflection[i] {
			return false
		}
	}
	return true
}

func main() {
	nubm := ""
	fmt.Println("Укажите число для проверки на палиндром")
	fmt.Scanln(&nubm)
	fmt.Println(palindromeSearch(nubm))
}
