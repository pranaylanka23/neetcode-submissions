func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2)<len(nums1){nums1,nums2 = nums2, nums1}
	total := len(nums1)+len(nums2)
	l,r := 0,len(nums1)
	for l<=r{
		i := l+(r-l)/2
		j := (total+1)/2-i
		lefta := math.MinInt64
		if i>0 { lefta = nums1[i-1]}
		leftb := math.MinInt64
		if j>0 { leftb = nums2[j-1]}
		righta := math.MaxInt64
		if i<len(nums1) { righta = nums1[i]}
		rightb := math.MaxInt64
		if j<len(nums2) { rightb = nums2[j]}
		if lefta<=rightb && leftb<=righta{
			if total%2==0{
				return (float64(max(lefta,leftb))+float64(min(rightb,righta)))/2.0
			}
			return float64(max(lefta,leftb))
		}else if lefta>rightb{
			r = i-1
		}else{
			l = i+1
		}
	}
	return -1
}
