package hashMap

func intersection(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return ret
	}

	if len(nums2) < len(nums1) {
		return intersectionCore(nums2, nums1)
	} else {
		return intersectionCore(nums1, nums2)
	}
}

func intersectionCore(nums, nums2 []int) []int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v] = 0
	}
	ret := make([]int, 0)
	for _, v := range nums2 {
		if _, ok := mp[v]; ok {
			mp[v] ++
		}
	}
	for k, _ := range mp {
		if mp[k] >= 1 {
			ret = append(ret, k)
		}
	}
	return ret
}
