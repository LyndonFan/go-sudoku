package main

import "fmt"

const N = 9

type Sudoku [N][N]int
type PossibleSudoku [N * N][N]bool

func (s Sudoku) Print() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if s[i][j] > 0 {
				fmt.Print(s[i][j])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (s Sudoku) ToPossibleSudoku() PossibleSudoku {
	ps := PossibleSudoku{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			v := s[i][j]
			if v > 0 {
				ps[i*N+j][v-1] = true
			} else {
				for k := 0; k < N; k++ {
					ps[i*N+j][k] = true
				}
			}
		}
	}
	return ps
}

func (ps PossibleSudoku) ToSudoku() Sudoku {
	s := Sudoku{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			val := -1
			ct := 0
			for k := 0; k < N; k++ {
				if ps[i*N+j][k] {
					ct++
					val = k + 1
				}
			}
			if ct == 1 {
				s[i][j] = val
			}
		}
	}
	return s
}
