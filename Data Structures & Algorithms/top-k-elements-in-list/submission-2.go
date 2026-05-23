func topKFrequent(nums []int, k int) []int {
	visit := make(map[int]int)
	for _,num := range nums{
		visit[num]++
	}
	freq := make([][]int, len(nums)+1)
	for key,val := range visit{
		freq[val] = append(freq[val], key)
	}
	ans := []int{}
	for i:=len(freq)-1; i>0; i--{
		for _, num := range freq[i]{
			ans = append(ans, num)
			if len(ans)==k{ return ans}
		}
	}
	return ans
}
