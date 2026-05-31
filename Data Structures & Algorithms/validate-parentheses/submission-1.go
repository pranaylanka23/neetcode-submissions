func isValid(s string) bool {
	if len(s)%2==1 {return false}
	par := map[rune]rune{'}':'{',']':'[',')':'('}
	stack := []rune{}
	for _,val := range s{
		if val == '{' || val == '[' || val == '(' {
			stack = append(stack, val)
			continue
		}
		if len(stack) <1 {return false}
		if stack[len(stack)-1] != par[val] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	if len(stack)>0 {
		return false
	}
	return true
}
