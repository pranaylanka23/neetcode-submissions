func asteroidCollision(asteroids []int) []int {
	ans := []int{}
	for _,val := range asteroids{
		for len(ans)>0 && val<0 && ans[len(ans)-1]>0{
			diff := val + ans[len(ans)-1]
			if diff < 0{
				ans = ans[:len(ans)-1]
			} else if diff > 0{
				val = 0
			}else{
				val = 0
				ans = ans[:len(ans)-1]
			}
		}
		if val!=0{
			ans = append(ans, val)
		}
	}
	return ans
}