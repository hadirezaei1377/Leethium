package sodoco

import "fmt"

func main() {

	board := [][]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("No solution exists")
	}
}

func printBoard(board [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
}

func findEmpty(board [][]int) (int, int, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func isValid(board [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	startRow, startCol := row/3*3, col/3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

func solveSudoku(board [][]int) bool {
	row, col, found := findEmpty(board)
	if !found {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num

			if solveSudoku(board) {
				return true
			}

			board[row][col] = 0
		}
	}

	return false
}
