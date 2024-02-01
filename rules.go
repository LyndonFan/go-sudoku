package main

type Rule interface {
	Global() bool
	Check(*PossibleSudoku, int, int) bool
	Apply(*PossibleSudoku, int, int) (bool, [][2]int)
}
