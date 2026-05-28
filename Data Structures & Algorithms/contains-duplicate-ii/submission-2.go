func containsNearbyDuplicate(nums []int, k int) bool {
	visit := make(map[int]int)
	for i,num := range nums{
		if j,ok := visit[num]; ok && i-j <=k{
			return true
		}
		visit[num]=i
	}
	return false
}
