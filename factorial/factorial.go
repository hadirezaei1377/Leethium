package factorial

import "fmt"

func main() {
	var number int
	fmt.Println("please enter a number :")
	fmt.Println("the factorial is equal to:", factorial(number))
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
