func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	check := make(map[[26]int][]string)
	for i:=0; i<len(strs); i++{
		w := strs[i]
		temp := [26]int{}
		for j:=0; j<len(w); j++{
			temp[w[j]-'a']++
		}
		check[temp] = append(check[temp], w)
	}
	for _,v := range check{
		ans = append(ans, v)
	}
	return ans
}
