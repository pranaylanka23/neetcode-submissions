func longestCommonPrefix(strs []string) string {
    ans := strs[0]
	for i:=1; i<len(strs); i++{
		ans = ans[0:min(len(ans),len(strs[i]))]
		for ans!=strs[i][0:len(ans)]{
			ans = ans[0:len(ans)-1]
		}
	}
	return ans
}
