package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i < sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := 99901
	if isPrime(n) {
		fmt.Printf("%d is a prime number\n", n)
	} else {
		fmt.Printf("%d is not a prime number\n", n)
	}

}

/*поиск максимального делителя

 */
