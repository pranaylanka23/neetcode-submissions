func maxSlidingWindow(nums []int, k int) []int {
    ans := []int{}
	r := []int{}
	for i:=0; i<len(nums); i++{
		if len(r)>0 && r[0]<i-k+1{
			r = r[1:]
		}
		for len(r)>0 && nums[r[len(r)-1]]<nums[i]{
			r = r[:len(r)-1]
		}
		r = append(r, i)
		if i>=k-1{ ans = append(ans, nums[r[0]])}
	}
	return ans
}
