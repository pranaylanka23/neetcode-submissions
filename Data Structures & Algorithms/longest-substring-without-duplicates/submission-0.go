func lengthOfLongestSubstring(s string) int {
	visit := make(map[byte]int)
	low,ans := 0,0
	for i:=0; i<len(s); i++{
		if val, ok := visit[s[i]]; ok{
			low = max(low,val+1)
		}
		visit[s[i]]=i
		ans = max(ans, i-low+1)
	}
	return ans
}
