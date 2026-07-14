func decodeString(s string) string {
	stack1 := []string{}
	stack2 := []int{}
	curr := ""
	k := 0

	for _,c := range s{
		if c>='0'&& c<='9'{
			k = k*10 + int(c-'0')
		} else if c == '['{
			stack2 = append(stack2, k)
			stack1 = append(stack1, curr)
			curr = ""
			k=0
		} else if c == ']'{
			smtg := curr
			curr = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			count := stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
			for i:=0; i< count; i++{
				curr += smtg
			}
		}else {
			curr += string(c)
		}
	}

	return curr
}
