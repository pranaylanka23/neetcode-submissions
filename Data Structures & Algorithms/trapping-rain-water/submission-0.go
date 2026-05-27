func trap(height []int) int {
	l,r := 0, len(height)-1
	maxl, maxr := height[l],height[r]
	maxArea := 0
	for l<r{
		if maxl<maxr{
			l++
			maxl = max(maxl, height[l])
			maxArea += maxl-height[l]
		} else {
			r--
			maxr = max(maxr, height[r])
			maxArea += maxr-height[r]
		}
	}
	return maxArea
}
