func findDuplicate(nums []int) int {
    n1,n2 := 0,0
	for{
		n1 = nums[n1]
		n2 = nums[nums[n2]]
		if n1==n2{break}
	}
	n1=0
	for n1!=n2{
		n1 = nums[n1]
		n2 = nums[n2]
	}
	return n1
}
