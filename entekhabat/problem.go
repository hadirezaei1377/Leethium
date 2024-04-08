package entekhabat

import "fmt"

func findWinner(n int) int {
	if n == 2 {
		return 1
	}
	return (2 * findWinner(n/2))
}

func main() {
	var n int
	fmt.Scanln(&n)
	winner := findWinner(n)
	fmt.Println(winner)
}
