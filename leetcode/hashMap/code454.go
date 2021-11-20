package hashMap

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 || len(nums3) == 0 || len(nums4) == 0 {
		return 0
	}
	mp1 := make(map[int]int, 0)
	for _, v := range nums1 {
		for _, vv := range nums2 {
			mp1[v + vv] ++
		}
	}
	count := 0
	for _, v := range nums3 {
		for _, vv := range nums4 {
			if _, ok := mp1[-(v + vv)]; ok {
				count += mp1[-(v + vv)]
			}
		}
	}

	return count
}

//func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
//	if len(nums1) == 0 || len(nums2) == 0 || len(nums3) == 0 || len(nums4) == 0 {
//		return 0
//	}
//	mp1 := make(map[int]int, 0)
//	mp2 := make(map[int]int, 0)
//	for _, v := range nums1 {
//		for _, vv := range nums2 {
//			mp1[v + vv] ++
//		}
//	}
//
//	for _, v := range nums3 {
//		for _, vv := range nums4 {
//			mp2[v + vv] ++
//		}
//	}
//
//	count := 0
//	for v, _ := range mp1 {
//		if _, ok := mp2[-v]; ok {
//			count += mp1[v] * mp2[-v]
//		}
//	}
//
//	return count
//}
