package main

import "fmt"

type Solver struct {
	possibilities []*PossibleSudoku
	Rules         []Rule
}

func NewSolver(rules []Rule) *Solver {
	return &Solver{
		possibilities: make([]*PossibleSudoku, 0, N*N),
		Rules:         rules,
	}
}

func (solver *Solver) Solve(puzzle Sudoku) []*Sudoku {
	ps := puzzle.ToPossibleSudoku()
	solver.possibilities = append(solver.possibilities, ps)
	var solved [N * N]bool
	var checkPosition [N * N]bool
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if puzzle[row][col] > 0 {
				solved[row*N+col] = true
				checkPosition[row*N+col] = true
			}
		}
	}
	var changed bool
	for _, b := range checkPosition {
		if b {
			changed = true
			break
		}
	}
	solutions := make([]*Sudoku, 0, N*N)
	for changed {
		changed = false
		var newCheckPosition [N * N]bool
		for pos, b := range checkPosition {
			if !b {
				continue
			}
			row, col := pos/N, pos%N
			for _, rule := range solver.Rules {
				satisfyRule := rule.Check(ps, row, col)
				if !satisfyRule {
					return []*Sudoku{}
				}
				applied, changedPositions := rule.Apply(ps, row, col)
				if !applied {
					continue
				}
				changed = true
				for _, cPos := range changedPositions {
					newCheckPosition[cPos[0]*N+cPos[1]] = true
				}
			}
		}
		checkPosition = newCheckPosition
	}
	ps.PrintWithSymbols()
	if ps.AllSolved() {
		solutions = append(solutions, ps.ToSudoku())
	} else {
		if ps.AllPossible() {
			fmt.Println("Multiple solutions, TODO")
		} else {
			fmt.Println("No solutions")
		}
	}
	return solutions
}
