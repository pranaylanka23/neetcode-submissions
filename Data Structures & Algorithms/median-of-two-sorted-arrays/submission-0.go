func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	totallen := len(merged)
	if totallen%2==0{
		return float64(merged[totallen/2-1]+merged[totallen/2])/2.0
	}
	return float64(merged[totallen/2])
}
