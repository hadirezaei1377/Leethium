package timeComplexity

import (
	"fmt"
)

// O(1)
func addElement(slice []int, element int) []int {
	return append(slice, element)
}

// O(n)
func removeElement(slice []int, element int) []int {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// O(n)
func searchElement(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func main() {
	slice := []int{10, 20, 30, 40, 50}

	slice = addElement(slice, 60)
	fmt.Println("List after adding 60:", slice)

	slice = removeElement(slice, 30)
	fmt.Println("List after removing 30:", slice)

	found := searchElement(slice, 20)
	fmt.Println("Is 20 in the list?", found)
}
