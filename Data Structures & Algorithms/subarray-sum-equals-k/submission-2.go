func subarraySum(nums []int, k int) int {
	prefix := map[int]int{0:1}
	ans, curr := 0,0
	for _, val := range nums{
		curr += val
		diff := curr - k
		ans += prefix[diff]
		prefix[curr]++
	}
	return ans
}
