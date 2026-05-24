func firstMissingPositive(nums []int) int {
	i,n := 0, len(nums)
	for i<n{
		if nums[i]<=0 || nums[i]>n{
			i++
			continue
		}
		ind := nums[i]-1
		if nums[ind]!=nums[i]{
			nums[ind],nums[i]=nums[i],nums[ind]
		} else{
			i++
		}
	}
	for i:=0; i<n; i++{
		if nums[i]!=i+1{
			return i+1
		}
	}
	return n+1
}
