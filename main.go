package main

func main() {
	s := Sudoku{
		{0, 0, 7, 0, 0, 5, 0, 0, 0},
		{0, 2, 0, 7, 3, 4, 0, 9, 0},
		{0, 0, 3, 0, 0, 0, 0, 0, 0},
		{8, 0, 9, 0, 7, 3, 0, 5, 0},
		{0, 0, 0, 8, 5, 0, 0, 0, 9},
		{3, 5, 0, 9, 0, 6, 0, 0, 0},
		{0, 3, 0, 0, 0, 0, 4, 0, 0},
		{6, 0, 8, 0, 4, 2, 0, 1, 0},
		{0, 0, 5, 0, 9, 0, 0, 8, 0},
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
