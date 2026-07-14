func simplifyPath(path string) string {
	stack := []string{}
	curr := ""
	for _, c := range path + "/" {
		if c =='/'{
			if curr == ".."{
				if len(stack)>0{
					stack = stack[:len(stack)-1]
				}
			} else if curr != "" && curr != "."{
				stack = append(stack,curr)
			}
			curr = ""
		}else {
			curr += string(c)
		}
	}

	return "/"+ strings.Join(stack, "/")
}
