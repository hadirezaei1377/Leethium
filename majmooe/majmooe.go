package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)

	input = strings.ReplaceAll(input, "}}", "}", -1)
	input = strings.ReplaceAll(input, "{{", "{", -1)

	total, nestedTotal := calculateSum(input)
	fmt.Println(total)
	fmt.Println(nestedTotal)
}

func calculateSum(input string) (int, int) {
	var total, nestedTotal int

	input = strings.TrimSpace(input[1 : len(input)-1])

	for _, item := range strings.Split(input, ",") {
		item = strings.TrimSpace(item)

		if isNumber(item) {
			num, _ := strconv.Atoi(item)
			total += num
			nestedTotal += num
		} else {

			nestedSum, nestedNestedTotal := calculateSum(item)
			total += nestedSum
			nestedTotal += nestedNestedTotal
		}
	}

	return total, nestedTotal
}

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
