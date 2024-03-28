package ghotrhaylozi

import (
	"fmt"
	"strings"
)

func printDiamond(n int) {

	for i := 0; i < n; i++ {
		spaces := strings.Repeat(" ", n-i-1)
		stars := strings.Repeat("*", 2*i+1)
		fmt.Printf("%s%s%s\n", spaces, stars, spaces)
	}

	for i := n - 2; i >= 0; i-- {
		spaces := strings.Repeat(" ", n-i-1)
		stars := strings.Repeat("*", 2*i+1)
		fmt.Printf("%s%s%s\n", spaces, stars, spaces)
	}
}

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scanln(&n)
	printDiamond(n)
}

// output for n= 5:
//     *
//    ***
//   *****
//  *******
// *********
//  *******
//   *****
//    ***
//     *
