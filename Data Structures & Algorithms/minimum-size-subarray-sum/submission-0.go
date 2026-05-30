func minSubArrayLen(target int, nums []int) int {
	ans,sum := math.MaxInt32,0
	l,r := 0,0
	for r<len(nums){
		sum += nums[r]
		for sum>=target{
			ans = min(ans,r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}
	if ans == math.MaxInt32 { return 0}
	return ans
}
