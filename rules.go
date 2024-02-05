package main

type Rule interface {
	Check(*PossibleSudoku, int, int) bool
	Apply(*PossibleSudoku, int, int) [][2]int
}
