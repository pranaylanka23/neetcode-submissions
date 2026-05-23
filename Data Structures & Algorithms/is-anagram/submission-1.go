func isAnagram(s string, t string) bool {
	if len(s) != len(t) {return false}
    ana1,ana2 := [26]int{},[26]int{}
	for i:=0; i<len(s); i++{
		ana1[s[i]-'a']++
		ana2[t[i]-'a']++
	}
	return ana1 == ana2
}
