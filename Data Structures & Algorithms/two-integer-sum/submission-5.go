func twoSum(nums []int, target int) []int {
    visit := make(map[int]int)
	for i:=0; i<len(nums); i++{
		if a,ok := visit[target-nums[i]]; ok{
			return []int{a,i}
		}
		visit[nums[i]]=i
	}
	return []int{}
}
