func groupAnagrams(strs []string) [][]string {
	ana := make(map[[26]int][]string)
	ans := [][]string{}
	for _,s := range strs{
		v := [26]int{}
		for _,val := range s{
			v[val-'a']++
		}
		ana[v] = append(ana[v], s)
	}
	for _,val := range ana{
		ans = append(ans, val)
	}
	return ans
}
