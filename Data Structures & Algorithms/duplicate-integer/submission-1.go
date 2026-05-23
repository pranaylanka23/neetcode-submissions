func hasDuplicate(nums []int) bool {
    val := make(map[int]bool)
	for _,num := range nums{
		if val[num] {return true}
		val[num] = true
	}
	return false
}
