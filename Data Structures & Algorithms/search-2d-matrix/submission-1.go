func searchMatrix(matrix [][]int, target int) bool {
	rows,cols := len(matrix), len(matrix[0])
	l,r := 0, (rows*cols)-1
	for l<=r {
		mid := l+ (r-l)/2
		val := matrix[mid/cols][mid%cols]
		if target==val{
			return true
		} else if target > val{
			l=mid+1
		} else{
			r= mid-1
		}
	}
	return false
}
