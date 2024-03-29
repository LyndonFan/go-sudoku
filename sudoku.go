package main

import (
	"fmt"
	"strings"
)

type Sudoku [9][9]int
type PossibleSudoku [9 * 9][9]bool

func (s Sudoku) Print() {
	horizontal := "┼---┼---┼---┼"
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Println(horizontal)
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				fmt.Print("|")
			}
			if s[i][j] > 0 {
				fmt.Print(s[i][j])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println(horizontal)
}

func (ps *PossibleSudoku) Get(row, col int) *[9]bool {
	return &(ps[row*9+col])
}

func (ps *PossibleSudoku) SetPossible(row, col, val int, possible bool) {
	ps[row*9+col][val] = possible
}

func (ps *PossibleSudoku) SetValue(row, col, val int) {
	for i := 0; i < 9; i++ {
		ps[row*9+col][i] = false
	}
	ps[row*9+col][val-1] = true
}

func (ps *PossibleSudoku) Possible(row, col int) bool {
	var ct int
	xs := ps.Get(row, col)
	for _, b := range xs {
		if b {
			ct++
		}
	}
	return ct > 0
}

func (ps *PossibleSudoku) Solved(row, col int) (bool, int) {
	val, ct := -1, 0
	xs := ps.Get(row, col)
	for k, b := range xs {
		if b {
			ct++
			val = k
		}
	}
	if ct == 1 {
		return true, val
	}
	return false, 0
}

func (ps *PossibleSudoku) AllPossible() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !ps.Possible(i, j) {
				return false
			}
		}
	}
	return true
}

func (ps *PossibleSudoku) AllSolved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			solved, _ := ps.Solved(i, j)
			if !solved {
				return false
			}
		}
	}
	return true
}

// Use trigram symbols (☰☱☲☳☴☵☶☷, 0x2630 - 0x2637)
// to show whether a cell can be set to a certian value
// breaks for yes, straight line for no
// use 3 such symbols to represent 1-9
// e.g. cell that can be 1, 5, 6, 9 would be ☱☲☶
// exception: cell has only one possible value
// in whihc case we just show the number
func (ps *PossibleSudoku) PrintWithSymbols() {
	symbols := [8]string{"☰", "☱", "☲", "☳", "☴", "☵", "☶", "☷"}
	horizontal := strings.Repeat("┼"+strings.Repeat("-", 11), 3) + "┼"
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Println(horizontal)
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			solved, val := ps.Solved(i, j)
			if solved {
				fmt.Printf(" %d ", val+1)
				continue
			}
			xs := ps.Get(i, j)
			toggles := [3]int{0, 0, 0}
			for k, b := range xs {
				if b {
					toggles[k%3] |= 1 << (k / 3)
				}
			}
			for _, v := range toggles {
				fmt.Print(symbols[v])
			}
		}
		fmt.Println("|")
	}
	fmt.Println(horizontal)
}

func (s Sudoku) ToPossibleSudoku() *PossibleSudoku {
	ps := PossibleSudoku{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := s[i][j]
			if v > 0 {
				ps[i*9+j][v-1] = true
			} else {
				for k := 0; k < 9; k++ {
					ps[i*9+j][k] = true
				}
			}
		}
	}
	return &ps
}

func (ps *PossibleSudoku) ToSudoku() *Sudoku {
	s := Sudoku{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := -1
			ct := 0
			for k := 0; k < 9; k++ {
				if ps[i*9+j][k] {
					ct++
					val = k + 1
				}
			}
			if ct == 1 {
				s[i][j] = val
			}
		}
	}
	return &s
}
