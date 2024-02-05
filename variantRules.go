package main

func checkNoRepeatWith(ps *PossibleSudoku, neighbourhood [][2]int, row, col int) bool {
	solved, val := ps.Solved(row, col)
	if !solved {
		return true
	}
	for _, pos := range neighbourhood {
		if pos[0] == row && pos[1] == col {
			continue
		}
		if ps.Get(pos[0], pos[1])[val] {
			return false
		}
	}
	return true
}

type DiagonalRule struct{}

func (rule DiagonalRule) neighbourhood(row, col int) [][2]int {
	if row != col && row+col != 8 {
		return nil
	}
	if row == 4 && col == 4 {
		res := make([][2]int, 17)
		for i := 0; i < 9; i++ {
			res[i] = [2]int{i, i}
		}
		for i := 0; i < 9; i++ {
			res[i+9] = [2]int{i, 8 - i}
		}
		return res
	}
	res := make([][2]int, 9)
	for i := 0; i < 9; i++ {
		res[i] = [2]int{i, i}
	}
	return res
}

func (rule DiagonalRule) Check(ps *PossibleSudoku, row, col int) bool {
	neighbourhood := rule.neighbourhood(row, col)
	if len(neighbourhood) == 0 {
		return true
	}
	return checkNoRepeat(ps, neighbourhood)
}

func (rule DiagonalRule) Apply(ps *PossibleSudoku, row, col int) [][2]int {
	neighbourhood := rule.neighbourhood(row, col)
	if len(neighbourhood) == 0 {
		return nil
	}
	return applyOneToNine(ps, neighbourhood, row, col)
}

type ChessKingRule struct{}

func (rule ChessKingRule) neighbourhood(row, col int) [][2]int {
	res := make([][2]int, 0, 9)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if row+i < 0 || row+i >= 9 || col+j < 0 || col+j >= 9 {
				continue
			}
			res = append(res, [2]int{row + i, col + j})
		}
	}
	return res
}

func (rule ChessKingRule) Check(ps *PossibleSudoku, row, col int) bool {
	neighbourhood := rule.neighbourhood(row, col)
	return checkNoRepeatWith(ps, neighbourhood, row, col)
}

func (rule ChessKingRule) Apply(ps *PossibleSudoku, row, col int) [][2]int {
	neighbourhood := rule.neighbourhood(row, col)
	return applyNoRepeat(ps, neighbourhood, row, col)
}

type ChessKnightRule struct{}

var chessKnightOffsets = [8][2]int{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}}

func (rule ChessKnightRule) neighbourhood(row, col int) [][2]int {
	res := make([][2]int, 0, 8)
	for _, offset := range chessKnightOffsets {
		newRow, newCol := row+offset[0], col+offset[1]
		if newRow < 0 || newRow >= 9 || newCol < 0 || newCol >= 9 {
			continue
		}
		res = append(res, [2]int{newRow, newCol})
	}
	return res
}

func (rule ChessKnightRule) Check(ps *PossibleSudoku, row, col int) bool {
	neighbourhood := rule.neighbourhood(row, col)
	return checkNoRepeatWith(ps, neighbourhood, row, col)
}

func (rule ChessKnightRule) Apply(ps *PossibleSudoku, row, col int) [][2]int {
	neighbourhood := rule.neighbourhood(row, col)
	return applyNoRepeat(ps, neighbourhood, row, col)
}
