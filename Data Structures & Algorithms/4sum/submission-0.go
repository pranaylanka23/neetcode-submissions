func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}

	for i:=0; i<len(nums); i++{
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		for j:=i+1; j<len(nums); j++{
			if j>i+1 && nums[j]==nums[j-1]{
				continue
			}
			l,r := j+1, len(nums)-1
			for l<r{
				sum := nums[i]+nums[j]+nums[l]+nums[r]
				if sum==target{
					ans = append(ans, []int{nums[i],nums[j],nums[l],nums[r]})
					l++
					r--
					for l<r && nums[l]==nums[l-1]{
						l++
					}
					for l<r && nums[r]==nums[r+1]{
						r--
					}
				} else if sum>target{
					r--
				}else{
					l++
				}
			}
		}
	}
	return ans
}
