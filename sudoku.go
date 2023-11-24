package sudoku

import (
	"fmt"
)

// based on a solution found here
// https://www.geeksforgeeks.org/sudoku-backtracking-7/

const (
	N = 9
	M = 3
)

type Board [N][N]int

// isSafe checks whether it will be
// legal to assign num to the
// given row, col
func (b *Board) isSafe(row, col, num int) bool {
	// Check if we find the same num
	// in the similar row or column, we
	// return false
	for i := 0; i < N; i++ {
		if b[row][i] == num || b[i][col] == num {
			return false
		}
	}

	// Check if we find the same num in
	// the particular 3*3 matrix,
	// we return false
	startRow := row - row%M
	startCol := col - col%M
	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			if b[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

// solveSudoku solve a partially filled-in grid and attempts
// to assign values to all unassigned locations in
// such a way to meet the requirements for
// Sudoku solution (non-duplication across rows,
// columns, and boxes)
func (b *Board) SolveSudoku(row, col int) bool {
	// Check if we have reached the 8th
	// row and 8th column (0
	// indexed matrix) , we are
	// returning true to avoid
	// further backtracking
	if row == N-1 && col == N-1 {
		return true
	}

	// Check if column value  becomes 9 ,
	// we move to next row and
	// column start from 0
	if col == N {
		row += 1
		col = 0
	}

	// Check if the current position of
	// the grid already contains
	// value >0, we iterate for next column
	if b[row][col] > 0 {
		return b.SolveSudoku(row, col+1)
	}

	for num := 1; num <= N; num++ {
		// Check if it is safe to place
		// the num (1-9)  in the
		// given row ,col  -> we
		// move to next column
		if b.isSafe(row, col, num) {
			// Assigning the num in
			// the current (row,col)
			// position of the grid
			// and assuming our assigned
			// num in the position
			// is correct
			b[row][col] = num
			// Checking for next possibility with next
			// column
			if b.SolveSudoku(row, col+1) {
				return true
			}
		}
		// Removing the assigned num ,
		// since our assumption
		// was wrong , and we go for
		// next assumption with
		// diff num value
		b[row][col] = 0
	}

	return false
}

// printing a utility function to print grid
func (b *Board) Printing() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(b[i][j], " ")
		}
		fmt.Println()
	}
}
