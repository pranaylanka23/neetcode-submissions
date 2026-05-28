func containsNearbyDuplicate(nums []int, k int) bool {
	for k>0{
		for i:=0; i<len(nums)-k; i++{
			if nums[i]==nums[i+k]{return true}
		}
		k--
	}
	return false
}
