func validPalindrome(s string) bool {
	isPalindrome := func(a,b int) bool{
		for a<b{
			if s[a]!=s[b]{return false}
			a++; b--
		}
		return true
	}
	l,r := 0, len(s)-1
	for l<r {
		if s[l]!=s[r]{
			return isPalindrome(l+1,r) || isPalindrome(l,r-1)
		}
		l++
		r--
	}
	return true
}