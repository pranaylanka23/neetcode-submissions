func twoSum(nums []int, target int) []int {
    visit := make(map[int]int)
	for i,num := range nums{
		visit[num] = i
	}
	for i, num := range nums{
		if val,ok := visit[target-num]; ok && val!=i{
			return []int{i,val}
		}
	}
	return []int{}
}
