package morraba

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter two positive integers:")
	fmt.Scan(&a, &b)

	if (a-b)%2 != 0 {
		fmt.Println("Wrong difference!")
		return
	}

	if a <= b {
		fmt.Println("Wrong order!")
		return
	}

	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			if i < (a-b)/2 || i >= a-(a-b)/2 || j < (a-b)/2 || j >= a-(a-b)/2 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
