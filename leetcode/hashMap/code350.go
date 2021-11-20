package hashMap

func intersect(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return ret
	}
	if len(nums1) < len(nums2) {
		return intersectCore(nums1, nums2)
	} else {
		return intersectCore(nums2, nums1)
	}
}

func intersectCore(numsMin, numsBigger []int) []int {
	mp := make(map[int]int)
	for _, v := range numsMin {
		mp[v] ++
	}
	ret := make([]int, 0)
	for _, v := range numsBigger {
		if _, ok := mp[v]; ok && mp[v] > 0 {
			ret = append(ret, v)
			mp[v] --
		}
	}
	return ret
}