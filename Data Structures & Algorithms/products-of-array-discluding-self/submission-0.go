func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i,_ := range nums {ans[i]=1}
	prefix := 1
	for i:=0;i<len(ans); i++{
		ans[i]= prefix
		prefix *= nums[i]
	}
	postfix := 1
	for i:=len(ans)-1; i>=0; i--{
		ans[i] *= postfix
		postfix *= nums[i]
	} 
	return ans
}