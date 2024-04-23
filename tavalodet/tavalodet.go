package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)

	year, _ := strconv.Atoi(input[:2])
	month, _ := strconv.Atoi(input[2:])

	persianYear := year + 621

	fmt.Printf("Sal: %04d\n", persianYear)
	fmt.Printf("Mah: %02d\n", month)
}
