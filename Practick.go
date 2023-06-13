package main

import "fmt"

func palindromeSearch(nubm string) {
	essence := []string{}
	reflection := []string{}
	for i := 0; i <= len(string(nubm))-1; i++ {
		essence = append(essence, string(nubm[i]))
	}
	for j := 1; j <= len(string(nubm)); j++ {
		reflection = append(reflection, string(nubm[len(string(nubm))-j]))
	}

	comparisonEssAndRefl(essence, reflection)
}

func comparisonEssAndRefl(ess, refl []string) {
	for i := 0; i < len(ess)-1; i++ {
		if ess[i] != refl[i] {
			fmt.Printf("Ваше число %s не является палиндромом\n", ess)
			return
		} else {
			fmt.Printf("Ваше число %s является палиндромом \n", ess)
			return
		}
	}

}
func main() {
	nubm := ""
	fmt.Println("Укажите число для проверки на палиндром")
	fmt.Scanln(&nubm)
	if len(string(nubm)) == 1 || nubm <= string("1") {
		fmt.Println("Палиндром является число которое больше 9 и не является отрицательным числом")
		return
	}
	palindromeSearch(nubm)
}
