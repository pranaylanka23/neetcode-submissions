func hasDuplicate(nums []int) bool {
    visit := make(map[int]bool)
	for i:=0; i<len(nums); i++{
		if _,ok := visit[nums[i]]; ok{
			return true
		}
		visit[nums[i]]=true
	}
	return false
}
