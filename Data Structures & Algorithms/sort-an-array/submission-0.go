func sortArray(nums []int) []int {
	check := true
	for check{
		check = false
		for i:=1; i<len(nums); i++{
			if nums[i]<nums[i-1]{
				nums[i],nums[i-1]=nums[i-1],nums[i]
				check = true
			}
		}
	}
	return nums
}
