func characterReplacement(s string, k int) int {
	l,maxf := 0,0
	count := make(map[byte]int)
	for i:=0; i<len(s); i++{
		count[s[i]]++
		maxf = max(maxf, count[s[i]])
		if (i-l+1)-maxf>k{
			count[s[l]]--
			l++
		}
	}
	return len(s)-l
}
