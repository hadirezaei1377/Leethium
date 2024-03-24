package tookhali

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the value of n:")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

// output for n = 3
// * * *
// *   *
// * * *
