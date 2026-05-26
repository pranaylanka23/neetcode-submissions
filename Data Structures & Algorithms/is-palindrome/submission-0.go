func isPalindrome(s string) bool {
	l,r := 0, len(s)-1
	for l<r{
		for l<r && !isValid(rune(s[l])){
			l++
		}
		for l<r && !isValid(rune(s[r])){
			r--
		}
		if unicode.ToLower(rune(s[l]))!=unicode.ToLower(rune(s[r])){
			return false
		}
		l++
		r--
	}
	return true
}

func isValid(b rune) bool{
	return unicode.IsDigit(b) || unicode.IsLetter(b)
}