func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i:=0; i<len(nums); i++{
		if nums[i]>0 {break}
		if i>0 && nums[i]==nums[i-1]{continue}
		l,r := i+1, len(nums)-1
		for l<r{
			target := nums[l]+nums[r]+nums[i]
			if target==0{
				ans = append(ans, []int{nums[l],nums[i],nums[r]})
				l++
				r--
				for l<r && nums[l]==nums[l-1]{l++}
			} else if target<0{
				l++
			} else{
				r--
			}
		}
	}
	return ans
}
