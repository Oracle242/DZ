package main

import "fmt"

// Архивация
func sorting(sumols string) {
	r := 1
	for i := 0; i < len(sumols)-1; i++ {
		if sumols[i] == sumols[i+1] {
			r++
		} else {
			fmt.Printf("%v%v", string(sumols[i]), r)
			r = 1
		}

	}
	fmt.Printf("%v%v", string(sumols[len(sumols)-1]), r)
	fmt.Printf("\n")
}

func main() {
	sumols := "zaaabbzbbccccccaaaaa"
	fmt.Println("Укажите строку для архивации")
	sorting(sumols)
}
