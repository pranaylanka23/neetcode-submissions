func minWindow(s string, t string) string {
	if len(t)>len(s) { return "" }
    count1 := make(map[byte]int)
	count2 := make(map[byte]int)
	for i:=0; i<len(t); i++{
		count2[t[i]]++
	}
	l,r,formed, minlen, minstart := 0,0,0,len(s)+1,0
	for r<len(s){
		count1[s[r]]++
		if count1[s[r]]==count2[s[r]]{ formed++ }
		for formed == len(count2) {
			if r-l+1<minlen{
				minlen = r-l+1
				minstart = l
			}
			count1[s[l]]--
			if count1[s[l]]<count2[s[l]]{
				formed--
			}
			l++
		}
		r++
	}
	if minlen == len(s)+1 {return ""}
	return s[minstart:minstart+minlen]
}
