func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens{
		switch token{
			case "+":
				a,b := stack[len(stack)-1],stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, a+b)
			case "-":
				a,b := stack[len(stack)-1],stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, b-a)
			case "*":
				a,b := stack[len(stack)-1],stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, a*b)
			case "/":
				a,b := stack[len(stack)-1],stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, b/a)
			default:
				num,_ := strconv.Atoi(token)
				stack = append(stack, num)
		}
	}
	return stack[0]
}
