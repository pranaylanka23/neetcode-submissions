func isAnagram(s string, t string) bool {
	if len(s)!=len(t) { return false}
	al := make([]int, 26)
	for i:=0; i<len(s); i++{
		al[s[i]-'a']++
		al[t[i]-'a']--
	}
	for i:=0; i<26; i++{
		if al[i]!=0{
			return false
		}
	}
	return true
}
