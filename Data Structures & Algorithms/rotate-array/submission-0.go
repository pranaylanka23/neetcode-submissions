func rotate(nums []int, k int) {
	k %= len(nums)
	ans := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums,ans)
}
