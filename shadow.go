package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	a, b := 30, 42
	fmt.Printf("GCD of %d and %d is %d\n", a, b, gcd(a, b))

}

/*поиск максимального делителя

 */
