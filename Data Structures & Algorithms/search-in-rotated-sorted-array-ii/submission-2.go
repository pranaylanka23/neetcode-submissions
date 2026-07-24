func search(nums []int, target int) bool {
	l,r := 0, len(nums)-1
	for l<r{
		mid := l + (r-l)/2
		if nums[mid]==target{ return true}
		if nums[l]<nums[mid]{
			if nums[l]<=target && target<=nums[mid]{
				r = mid
			}else{
				l = mid+1
			}
		}else if nums[l]>nums[mid]{
			if nums[mid]<=target && target<=nums[r]{
				l=mid
			}else{
				r=mid-1
			}
		}else{
			l++
		}
	}
	if nums[l]==target{return true}
	return false
}
