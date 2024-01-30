package main

func checkNoRepeat(ps PossibleSudoku, neighbourhood [][2]int) bool {
	var seen [N]bool
	var possible [N]bool
	for _, pos := range neighbourhood {
		solved, val := ps.Solved(pos[0], pos[1])
		if solved {
			if seen[val] {
				return false
			}
			seen[val] = true
			possible[val] = true
			continue
		}
		xs := ps.Get(pos[0], pos[1])
		for i, b := range xs {
			if b {
				possible[i] = true
			}
		}
	}
	for _, b := range possible {
		if !b {
			return false
		}
	}
	return true
}

func applyNoRepeat(
	ps PossibleSudoku,
	neighbourhood [][2]int,
	row, col int,
) (bool, [][2]int) {
	// below is naive approach, only does work if the cell is solved
	// TODO: do some complex rules? Or decide to leave it to another rule
	solved, val := ps.Solved(row, col)
	if !solved {
		return false, nil
	}
	res := make([][2]int, 0, N) // usually N, but can be smaller or larger
	for _, pos := range neighbourhood {
		if pos[0] == row && pos[1] == col {
			continue
		}
		xs := ps.Get(pos[0], pos[1])
		if xs[val] {
			res = append(res, pos)
			xs[val] = false
		}
	}
	return true, res
}

type HorizontalRule struct{}

func (rule HorizontalRule) neighbourhood(row, _ int) [][2]int {
	res := make([][2]int, 9)
	for i := 0; i < 9; i++ {
		res[i] = [2]int{row, i}
	}
	return res
}

func (rule HorizontalRule) Global() bool {
	return true
}

func (rule HorizontalRule) Check(ps PossibleSudoku, row, col int) bool {
	return checkNoRepeat(ps, rule.neighbourhood(row, col))
}

func (rule HorizontalRule) Apply(ps PossibleSudoku, row, col int) (bool, [][2]int) {
	return applyNoRepeat(ps, rule.neighbourhood(row, col), row, col)
}

type VerticalRule struct{}

func (rule VerticalRule) neighbourhood(_, col int) [][2]int {
	res := make([][2]int, 9)
	for i := 0; i < 9; i++ {
		res[i] = [2]int{i, col}
	}
	return res
}

func (rule VerticalRule) Global() bool {
	return true
}

func (rule VerticalRule) Check(ps PossibleSudoku, row, col int) bool {
	return checkNoRepeat(ps, rule.neighbourhood(row, col))
}

func (rule VerticalRule) Apply(ps PossibleSudoku, row, col int) (bool, [][2]int) {
	return applyNoRepeat(ps, rule.neighbourhood(row, col), row, col)
}

type SquareRule struct{}

func (rule SquareRule) Global() bool {
	return true
}

func (rule SquareRule) neighbourhood(row, col int) [][2]int {
	squareRow, squareCol := row/3, col/3
	res := make([][2]int, 9)
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			res[r*3+c] = [2]int{squareRow*3 + r, squareCol*3 + c}
		}
	}
	return res
}

func (rule SquareRule) Check(ps PossibleSudoku, row, col int) bool {
	return checkNoRepeat(ps, rule.neighbourhood(row, col))
}

func (rule SquareRule) Apply(ps PossibleSudoku, row, col int) (bool, [][2]int) {
	return applyNoRepeat(ps, rule.neighbourhood(row, col), row, col)
}
