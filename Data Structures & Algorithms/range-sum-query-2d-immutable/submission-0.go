type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows, cols := len(matrix), len(matrix[0])
	sumMat := make([][]int, rows+1)
	for i := range sumMat{
		sumMat[i]= make([]int, cols+1)
	}
	for r:=0; r<rows; r++{
		prefix :=0
		for c:=0; c<cols; c++{
			prefix += matrix[r][c]
			above := sumMat[r][c+1]
			sumMat[r+1][c+1]=prefix+above
		}
	}
	return NumMatrix{matrix: sumMat}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1++; col1++; row2++; col2++
	br := this.matrix[row2][col2]
	tl := this.matrix[row1-1][col1-1]
	tr := this.matrix[row1-1][col2]
	bl := this.matrix[row2][col1-1]
	return br -tr -bl + tl
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)