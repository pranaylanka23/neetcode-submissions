func checkInclusion(s1 string, s2 string) bool {
	if len(s2)<len(s1) {return false}
	a1, a2 := [26]int{}, [26]int{}
	for i,_ := range s1 {
		a1[s1[i]-'a']++
		a2[s2[i]-'a']++
	}
	if a1 == a2 { return true}
	for i:=len(s1); i<len(s2); i++{
		a2[s2[i-len(s1)]-'a']--
		a2[s2[i]-'a']++
		if a1==a2 { return true}
	}
	return false
}
