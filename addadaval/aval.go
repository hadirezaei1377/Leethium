package addadaval

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func printPrimesInRange(a, b int) {
	primeNumbers := []int{}
	for i := a + 1; i < b; i++ {
		if isPrime(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}
	fmt.Println("Prime numbers in the range", a, "to", b, "are:", primeNumbers)
}

func printAllPrimes(n int) {
	primeNumbers := []int{}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}
	fmt.Println("All prime numbers up to", n, "are:", primeNumbers)
}

func main() {
	var a, b int
	fmt.Println("Enter the range (a,b):")
	fmt.Scan(&a, &b)

	// Print prime numbers in the given range
	printPrimesInRange(a, b)

	// Print all prime numbers up to b
	printAllPrimes(b)
}
