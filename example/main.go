package main

import (
	"fmt"
	"time"

	"github.com/oussamarouabah/sudoku"
)

func main() {
	start := time.Now()
	solveSudoku()
	take := time.Since(start)
	fmt.Printf("is took %s to solve the problem\n", take.String())
}

func solveSudoku() {

	// 0 means unassigned cells
	var b = sudoku.Board{
		[sudoku.N]int{3, 0, 6, 5, 0, 8, 4, 0, 0},
		[sudoku.N]int{5, 2, 0, 0, 0, 0, 0, 0, 0},
		[sudoku.N]int{0, 8, 7, 0, 0, 0, 0, 3, 1},
		[sudoku.N]int{0, 0, 3, 0, 1, 0, 0, 8, 0},
		[sudoku.N]int{9, 0, 0, 8, 6, 3, 0, 0, 5},
		[sudoku.N]int{0, 5, 0, 0, 9, 0, 6, 0, 0},
		[sudoku.N]int{1, 3, 0, 0, 0, 0, 2, 5, 0},
		[sudoku.N]int{0, 0, 0, 0, 0, 0, 0, 7, 4},
		[sudoku.N]int{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	if b.SolveSudoku(0, 0) {
		b.Printing()
	} else {
		fmt.Println("no solution  exists ")
	}
}
