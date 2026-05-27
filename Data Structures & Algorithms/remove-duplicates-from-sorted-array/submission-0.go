func removeDuplicates(nums []int) int {
	l,r := 0,0
	count:=1
	for r<len(nums){
		if nums[l]==nums[r]{
			r++
		}else{
			l++
			nums[l]=nums[r]
			count++
		}
	}
	return count
}
