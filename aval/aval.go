package aval

import (
	"fmt"
	"math"
)

func IsPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func PrimeNumbersBetween(start, end int) []int {
	var primeNumbers []int
	for num := start; num <= end; num++ {
		if IsPrime(num) {
			primeNumbers = append(primeNumbers, num)
		}
	}
	return primeNumbers
}

func main() {
	var startNum, endNum int
	fmt.Print("Enter the First number: ")
	fmt.Scan(&startNum)
	fmt.Print("Enter the Last number: ")
	fmt.Scan(&endNum)

	primeNums := PrimeNumbersBetween(startNum, endNum)
	fmt.Printf("Prime numbers between %d and %d are: %v\n", startNum, endNum, primeNums)
}
