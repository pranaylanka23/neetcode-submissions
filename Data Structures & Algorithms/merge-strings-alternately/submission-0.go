func mergeAlternately(word1 string, word2 string) string {
	ans := ""
	small := min(len(word1),len(word2))
	for i:=0; i<small; i++{
		ans += string(word1[i])+ string(word2[i])
	}
	if len(word1)<len(word2) { 
		ans += word2[small:]
	} else {
		ans += word1[small:]
	}
	return ans
}
