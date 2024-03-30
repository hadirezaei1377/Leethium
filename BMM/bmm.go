package bmm

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	var a, b int
	fmt.Println("Enter two numbers:")
	fmt.Scanln(&a, &b)

	result := gcd(a, b)
	fmt.Printf("GCD of %d and %d is: %d\n", a, b, result)
}
