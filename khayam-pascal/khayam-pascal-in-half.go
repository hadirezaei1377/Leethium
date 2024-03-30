package khayampascal

import "fmt"

func main() {
	var numRows int
	fmt.Print("تعداد سطرهای مثلث خیام پاسکال را وارد کنید: ")
	fmt.Scanln(&numRows)

	pascalTriangle := make([][]int, numRows)
	for i := range pascalTriangle {
		pascalTriangle[i] = make([]int, i+1)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				pascalTriangle[i][j] = 1
			} else {
				pascalTriangle[i][j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
			}
		}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", pascalTriangle[i][j])
		}
		fmt.Println()
	}
}

// output for n=5

// 1
// 1 1
// 1 2 1
// 1 3 3 1
// 1 4 6 4 1
