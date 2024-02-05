package main

import "fmt"

type Solver struct {
	possibilities []*PossibleSudoku
	Rules         []Rule
}

func NewSolver(rules []Rule) *Solver {
	return &Solver{
		possibilities: make([]*PossibleSudoku, 0, 9*9),
		Rules:         rules,
	}
}

func (solver *Solver) Solve(puzzle Sudoku) []*Sudoku {
	ps := puzzle.ToPossibleSudoku()
	solver.possibilities = append(solver.possibilities, ps)
	var solved [9 * 9]bool
	var checkPosition [9 * 9]bool
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if puzzle[row][col] > 0 {
				solved[row*9+col] = true
				checkPosition[row*9+col] = true
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
	solutions := make([]*Sudoku, 0, 9*9)
	for changed {
		changed = false
		var newCheckPosition [9 * 9]bool
		for pos, b := range checkPosition {
			if !b {
				continue
			}
			row, col := pos/9, pos%9
			for _, rule := range solver.Rules {
				// satisfyRule := rule.Check(ps, row, col)
				// if !satisfyRule {
				// 	return []*Sudoku{}
				// }
				changedPositions := rule.Apply(ps, row, col)
				for _, cPos := range changedPositions {
					newCheckPosition[cPos[0]*9+cPos[1]] = true
				}
				changed = changed || (len(changedPositions) > 0)
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
