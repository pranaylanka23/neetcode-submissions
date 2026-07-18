func mySqrt(x int) int {
	l,r := 0, x
	ans := 0
	for l<=r{
		m := l + (r-l)/2
		if int64(m)*int64(m) > int64(x) {
			r = m-1
		} else if int64(m)*int64(m) < int64(x) {
			l= m+1
			ans = m
		} else {
			return m
		}
	}
	return ans
}
