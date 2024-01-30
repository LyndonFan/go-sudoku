package main

import "fmt"

type Rule interface {
	Global() bool
	Check(PossibleSudoku, int, int) bool
	Apply(PossibleSudoku, int, int) (bool, [][2]int)
}

type Solver struct {
	possibilities []PossibleSudoku
	GlobalRules   []Rule
	LocalRules    []Rule
}

func NewSolver(rules []Rule) *Solver {
	globalRules := make([]Rule, 0, len(rules))
	localRules := make([]Rule, 0, len(rules))
	for _, rule := range rules {
		if rule.Global() {
			globalRules = append(globalRules, rule)
		} else {
			localRules = append(localRules, rule)
		}
	}
	return &Solver{
		possibilities: make([]PossibleSudoku, 0, N*N),
		GlobalRules:   globalRules,
		LocalRules:    localRules,
	}
}

func (solver *Solver) Solve(puzzle Sudoku) []Sudoku {
	ps := puzzle.ToPossibleSudoku()
	solver.possibilities = append(solver.possibilities, ps)
	var solved [N * N]bool
	positionQueue := make([]int, 0, N*N)
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if puzzle[row][col] > 0 {
				solved[row*N+col] = true
				positionQueue = append(positionQueue, row*N+col)
			}
		}
	}
	changed := len(positionQueue) > 0
	solutions := make([]Sudoku, 0, N*N)
	for changed {
		changed = false
		var addPosition [N * N]bool
		for _, pos := range positionQueue {
			row, col := pos/N, pos%N
			for _, rule := range solver.GlobalRules {
				satisfyRule := rule.Check(ps, row, col)
				if !satisfyRule {
					return []Sudoku{}
				}
				applied, changedPositions := rule.Apply(ps, row, col)
				if !applied {
					continue
				}
				changed = true
				for _, cPos := range changedPositions {
					addPosition[cPos[0]*N+cPos[1]] = true
				}
			}
		}
		positionQueue = positionQueue[:0]
		for i, b := range addPosition {
			if b {
				positionQueue = append(positionQueue, i)
			}
		}
	}
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
