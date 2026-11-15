func longestCommonPrefix(strs []string) string {
    ans := strs[0]
	for i:=1; i<len(strs); i++{
		curr := strs[i]
		if len(ans)<=len(strs[i]){
			curr = curr[:len(ans)]
		}else{
			ans = ans[:len(curr)]
		}
		for ans!=curr{
			curr = curr[:len(curr)-1]
			ans = ans[:len(ans)-1]
		}
	}
	return ans
}
