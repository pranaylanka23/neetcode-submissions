func findClosestElements(arr []int, k int, x int) []int {
	ans := []int{}
	for i:=0; i<len(arr); i++{
		if len(ans)<k {
			ans = append(ans, arr[i])
			continue
		}
		if abs(ans[0],x) > abs(arr[i],x){
			ans = ans[1:]
			ans = append(ans, arr[i])
		}
	}
	return ans
}

func abs(a,b int) int{
	if a-b < 0 { return b-a}
	return a-b
}
