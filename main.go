package main

func main() {
	s := Sudoku{
		{4, 0, 0, 0, 0, 9, 3, 0, 0},
		{0, 0, 0, 4, 1, 0, 0, 6, 0},
		{0, 0, 0, 5, 3, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 6, 5, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 7, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 9, 8, 0, 0, 0},
		{0, 7, 0, 0, 2, 4, 0, 0, 0},
		{0, 0, 9, 3, 0, 0, 0, 0, 4},
	}
	s.Print()

	ps := s.ToPossibleSudoku()
	ps.PrintWithSymbols()

	solver := NewSolver([]Rule{
		VerticalRule{},
		HorizontalRule{},
		SquareRule{},
		ChessKingRule{},
		ChessKnightRule{},
	})
	solutions := solver.Solve(s)
	for _, solution := range solutions {
		solution.Print()
	}
}
