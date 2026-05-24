func majorityElement(nums []int) []int {
	n1,n2,c1,c2 := -1,-1,0,0
	for _,val := range nums{
		if val==n1{
			c1++
		}else if val==n2{
			c2++
		}else if c1==0{
			c1++
			n1=val
		}else if c2==0{
			c2++
			n2=val
		}else{
			c1--
			c2--
		}
	}
	c1,c2 = 0,0
	for _,val := range nums{
		if val==n1{
			c1++
		}else if val==n2{
			c2++
		}
	}
	ans := []int{}
	if c1>len(nums)/3 {ans = append(ans, n1)}
	if c2>len(nums)/3 {ans = append(ans, n2)}
	return ans
}
