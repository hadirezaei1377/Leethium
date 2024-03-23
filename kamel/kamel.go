package kamel

import "fmt"

func main() {
	var num int
	fmt.Println("لطفاً یک عدد وارد کنید:")
	fmt.Scanln(&num)

	var slices []int

	for i := 1; i < num; i++ {
		if num%i == 0 {
			slices = append(slices, i)
		}
	}

	sum := 0
	for _, value := range slices {
		sum += value
	}

	if sum == num {
		fmt.Println("عدد کامل است")
	} else {
		fmt.Println("عدد کامل نیست")
	}
}
