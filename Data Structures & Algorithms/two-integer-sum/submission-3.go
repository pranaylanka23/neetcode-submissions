func twoSum(nums []int, target int) []int {
    visit := make(map[int]int)
	for i, num := range nums{
		diff := target - num
		if j, ok := visit[diff]; ok{
			return []int{j,i}
		}
		visit[num]=i
	}
	return []int{}
}
