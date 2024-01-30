package main

func main() {
	s := Sudoku{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 1, 5, 6, 4, 8, 9, 7},
		{5, 6, 4, 8, 9, 7, 2, 3, 1},
		{8, 9, 7, 2, 3, 1, 5, 6, 4},
		{3, 1, 2, 6, 4, 5, 9, 7, 8},
		{6, 4, 5, 9, 7, 8, 3, 1, 2},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	s.Print()

	ps := s.ToPossibleSudoku()
	ps.PrintWithSymbols()

	solver := NewSolver([]Rule{
		VerticalRule{},
		HorizontalRule{},
		SquareRule{},
	})
	solutions := solver.Solve(s)
	for _, solution := range solutions {
		solution.Print()
	}
}
