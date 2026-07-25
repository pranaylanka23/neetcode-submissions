func splitArray(nums []int, k int) int {
	cansplit := func(m int) bool{
		count,curr := 1,0
		for _,val := range nums{
			curr += val
			if curr>m{
				count++
				if count>k { return false}
				curr = val
			}
		}
		return true
	}
	l,r := 0, 0
	for _, num := range nums{
		l = max(l,num)
		r += num
	}
	ans := r
	for l<=r{
		m := l + (r-l)/2
		if cansplit(m){
			ans = m
			r= m-1
		}else{
			l=m+1
		}
	}
	return ans
}
