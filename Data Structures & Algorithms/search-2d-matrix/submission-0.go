func searchMatrix(matrix [][]int, target int) bool {
	rows,cols := len(matrix), len(matrix[0])
	l,r := 0, rows-1
	for l<=r{
		mid := l + (r-l)/2
		if matrix[mid][cols-1]<target{
			l = mid + 1
		} else if matrix[mid][0] > target {
			r=mid-1
		} else{
			break
		}
	}

	if l>r {return false}
	row := (l+r)/2
	l,r = 0, cols-1
	for l<=r{
		m := (l+r)/2
		if target>matrix[row][m]{
			l=m+1
		} else if target<matrix[row][m]{
			r = m-1
		}else{
			return true
		}
	}
	return false
}
