func minEatingSpeed(piles []int, h int) int {
	l,r := 1,0
	for _,p := range piles{
		r = max(r, p)
	}
	ans := r

	for l<=r{
		m := (l+r)/2
		time := 0

		for _,p := range piles{
			time += int(math.Ceil(float64(p)/float64(m)))
		}
		if time<=h{
			ans = m
			r = m-1
		} else {
			l=m+1
		}
	}
	return ans
}
