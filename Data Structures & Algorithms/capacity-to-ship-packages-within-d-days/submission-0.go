func shipWithinDays(weights []int, days int) int {
	l,r := 0,0
	for _, val := range weights{
		l = max(l,val)
		r+=val
	}
	ans  := r
	canShip := func(val int) bool{
		ships, curr := 1, val
		for _,v := range weights{
			if curr-v<0{
				ships++
				if ships>days{return false}
				curr = val
			}
			curr -= v
		}
		return true
	}

	for l<=r{
		mid := (l+r)/2
		if canShip(mid){
			ans = min(ans,mid)
			r = mid-1
		}else{
			l = mid+1
		}
	}
	return ans
}
