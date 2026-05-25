func maxArea(heights []int) int {
	maxArea := 0
	l,r := 0, len(heights)-1
	for l<r{
		area := min(heights[l],heights[r])*(r-l)
		maxArea = max(maxArea, area)
		if heights[l]<heights[r]{
			l++
		}else {
			r--
		}
	}
	return maxArea
}
