package main

import "fmt"

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

// ?dsf
