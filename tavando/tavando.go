package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	power := math.Log2(float64(n)) + 1
	result := int(math.Pow(2, math.Ceil(power)))

	fmt.Println(result)
}
