/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
    l,r,peak := 0, mountainArr.length()-1,0
	for l<=r{
		m := l + (r-l)/2
		left := mountainArr.get(m-1)
		mid := mountainArr.get(m)
		right := mountainArr.get(m+1)
		if left<mid && mid<right{
			l= m+1
		}else if left>mid && mid >right{
			r = m-1
		}else{
			peak = m
			break
		}
	}
	l,r = 0,peak
	for l<=r{
		mid := l + (r-l)/2
		if mountainArr.get(mid)==target{return mid}
		if mountainArr.get(mid)>target{
			r=mid-1
		} else{
			l=mid+1
		}
	}
	l,r = peak+1, mountainArr.length()-1
	for l<=r{
		mid := l+(r-l)/2
		if mountainArr.get(mid) == target { return mid}
		if mountainArr.get(mid)>target{
			l=mid+1
		}else{
			r=mid-1
		}
	}
	return -1
}
