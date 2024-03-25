package max

import (
	"fmt"
	"strconv"
)

func main() {
	var numbers []int
	var input string

	fmt.Println("Enter numbers separated by commas, type 'finish' when you aredone:")

	for {
		fmt.Scan(&input)
		if input == "finish" {
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number or 'finish' to end.")
			continue
		}
		numbers = append(numbers, number)
	}

	max := findMax(numbers)
	fmt.Printf("Max of %s is %d\n", numbers, max)
}

func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}
