package farzad

import "fmt"

func multiplyMatrices(A, B [][]int, n int) [][]int {
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, n)
		for j := range C[i] {
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

func determinant(matrix [][]int, n int) int {
	if n == 1 {
		return matrix[0][0]
	}

	sign := 1
	det := 0
	temp := make([][]int, n)
	for i := range temp {
		temp[i] = make([]int, n)
	}

	for f := 0; f < n; f++ {
		getCofactor(matrix, temp, 0, f, n)
		det += sign * matrix[0][f] * determinant(temp, n-1)
		sign = -sign
	}

	return det
}

func getCofactor(matrix, temp [][]int, p, q, n int) {
	i, j := 0, 0
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row != p && col != q {
				temp[i][j] = matrix[row][col]
				j++
				if j == n-1 {
					j = 0
					i++
				}
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	A := make([][]int, n)
	B := make([][]int, n)

	for i := range A {
		A[i] = make([]int, n)
		for j := range A[i] {
			fmt.Scan(&A[i][j])
		}
	}

	for i := range B {
		B[i] = make([]int, n)
		for j := range B[i] {
			fmt.Scan(&B[i][j])
		}
	}

	C := multiplyMatrices(A, B, n)
	det := determinant(C, n)

	if det%2 != 0 {
		fmt.Println("Danial")
	} else {
		fmt.Println("Farzad")
	}
}
