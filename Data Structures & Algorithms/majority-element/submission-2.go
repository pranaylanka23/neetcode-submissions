func majorityElement(nums []int) int {
    num, count := nums[0],0
	for _,val := range nums{
		if val==num { 
			count++ 
		} else {
			count--
		}
		if count==0 { 
			num = val
			count++
		}
	}
	return num
}
