func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	ans := make([]int, len(temperatures))
	for i,val := range temperatures{
		for len(stack)>0 && val> temperatures[stack[len(stack)-1]]{
			ans[stack[len(stack)-1]]=i-stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
