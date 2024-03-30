package khayampascal

// import "fmt"

// func main() {
// 	var numRows int
// 	fmt.Print("تعداد سطرهای مثلث خیام پاسکال را وارد کنید: ")
// 	fmt.Scanln(&numRows)

// 	pascalTriangle := make([][]int, numRows)
// 	for i := range pascalTriangle {
// 		pascalTriangle[i] = make([]int, 2*i+1)
// 	}

// 	for i := 0; i < numRows; i++ {
// 		for j := 0; j <= i; j++ {
// 			if j == 0 || j == i {
// 				pascalTriangle[i][i-j] = 1
// 			} else {
// 				pascalTriangle[i][i-j] = pascalTriangle[i-1][i-j-1] + pascalTriangle[i-1][i-j+1]
// 			}
// 		}
// 	}

// 	for i := 0; i < numRows; i++ {
// 		for j := 0; j < numRows-i; j++ {
// 			fmt.Print(" ")
// 		}
// 		for j := 0; j < 2*i+1; j++ {
// 			fmt.Printf("%d ", pascalTriangle[i][j])
// 		}
// 		fmt.Println()
// 	}
// }
