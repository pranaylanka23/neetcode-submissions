func isValidSudoku(board [][]byte) bool {
	row, col, box := [9][9]bool{},[9][9]bool{},[9][9]bool{}
	for i,r := range board{
		for j, num := range r{
			if num == '.' {continue}
			k := int(num)-49
			if row[i][k] || col[j][k] || box[(i/3)*3+(j/3)][k] { return false}
			row[i][k], col[j][k], box[(i/3)*3+(j/3)][k] = true, true, true
		}
	}
	return true
}