package fibonacci

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number of Fibonacci sentences you want to print:")
	fmt.Scan(&n)

	printFibonacciSentences(n)
}

func printFibonacciSentences(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Fibonacci sentence: %d\n", i+1, a)
		a, b = b, a+b
	}
}
