package main

func checkNoRepeat(ps *PossibleSudoku, neighbourhood [][2]int) bool {
	var seen [9]bool
	var possible [9]bool
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
	ps *PossibleSudoku,
	neighbourhood [][2]int,
	row, col int,
) [][2]int {
	// below is naive approach, only does work if the cell is solved
	// TODO: do some complex rules? Or decide to leave it to another rule
	solved, val := ps.Solved(row, col)
	if !solved {
		return nil
	}
	res := make([][2]int, 0, 9) // usually N, but can be smaller or larger
	for _, pos := range neighbourhood {
		if pos[0] == row && pos[1] == col {
			continue
		}
		if ps.Get(pos[0], pos[1])[val] {
			res = append(res, pos)
			ps.SetPossible(pos[0], pos[1], val, false)
		}
	}
	return res
}

func applyOneToNine(
	ps *PossibleSudoku,
	neighbourhood [][2]int,
	row, col int,
) [][2]int {
	solved, _ := ps.Solved(row, col)
	if solved {
		return applyNoRepeat(ps, neighbourhood, row, col)
	}
	neighbourIndex := -1
	for i, pos := range neighbourhood {
		if pos[0] == row && pos[1] == col {
			neighbourIndex = i
			break
		}
	}
	if neighbourIndex == -1 {
		return nil
	}
	possibleIndexes := make([][]int, 9)
	for i := 0; i < 9; i++ {
		possibleIndexes[i] = make([]int, 0, 9)
	}
	for idx, pos := range neighbourhood {
		xs := ps.Get(pos[0], pos[1])
		for i, b := range xs {
			if b {
				possibleIndexes[i] = append(possibleIndexes[i], idx)
			}
		}
	}
	for v, idxs := range possibleIndexes {
		if !(len(idxs) == 1 && idxs[0] == neighbourIndex) {
			continue
		}
		for x := 0; x < 9; x++ {
			ps.SetPossible(row, col, x, x == v)
		}
		return applyNoRepeat(ps, neighbourhood, row, col)
	}
	return nil
}

type HorizontalRule struct{}

func (rule HorizontalRule) neighbourhood(row, _ int) [][2]int {
	res := make([][2]int, 9)
	for i := 0; i < 9; i++ {
		res[i] = [2]int{row, i}
	}
	return res
}

func (rule HorizontalRule) Check(ps *PossibleSudoku, row, col int) bool {
	return checkNoRepeat(ps, rule.neighbourhood(row, col))
}

func (rule HorizontalRule) Apply(ps *PossibleSudoku, row, col int) [][2]int {
	return applyOneToNine(ps, rule.neighbourhood(row, col), row, col)
}

type VerticalRule struct{}

func (rule VerticalRule) neighbourhood(_, col int) [][2]int {
	res := make([][2]int, 9)
	for i := 0; i < 9; i++ {
		res[i] = [2]int{i, col}
	}
	return res
}

func (rule VerticalRule) Check(ps *PossibleSudoku, row, col int) bool {
	return checkNoRepeat(ps, rule.neighbourhood(row, col))
}

func (rule VerticalRule) Apply(ps *PossibleSudoku, row, col int) [][2]int {
	return applyOneToNine(ps, rule.neighbourhood(row, col), row, col)
}

type SquareRule struct{}

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

func (rule SquareRule) Check(ps *PossibleSudoku, row, col int) bool {
	return checkNoRepeat(ps, rule.neighbourhood(row, col))
}

func (rule SquareRule) Apply(ps *PossibleSudoku, row, col int) [][2]int {
	return applyOneToNine(ps, rule.neighbourhood(row, col), row, col)
}
