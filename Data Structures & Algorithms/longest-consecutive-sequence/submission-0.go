func longestConsecutive(nums []int) int {
	count := make(map[int]int)
	ans := 0
	for _,num := range nums{
		if count[num]==0{
			l,r := count[num-1], count[num+1]
			sum := l+r+1
			count[num],count[num-l],count[num+r]=sum,sum,sum
			ans =max(ans,sum)
		}
	}
	return ans
}