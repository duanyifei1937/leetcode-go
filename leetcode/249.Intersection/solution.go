package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	// 预先判断nil slice
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	nums1_map := make(map[int]int)

	// 反转：使得key唯一
	for i := 0; i < len(nums1); i++ {
		if n, ok := nums1_map[nums1[i]]; ok {
			nums1_map[nums1[i]] = n + 1
		} else {
			nums1_map[nums1[i]] = 1
		}
	}

	nums2_map := make(map[int]int)
	for i := range nums2 {
		if n, ok := nums2_map[nums2[i]]; ok {
			nums2_map[nums2[i]] = n + 1
		} else {
			nums2_map[nums2[i]] = 1
		}
	}

	// 根据len,从small len遍历
	if len(nums1_map) > len(nums2_map) {
		nums1_map, nums2_map = nums2_map, nums1_map
	}

	// 结果map
	result := make(map[int]int)
	for k := range nums1_map {
		_, ok := nums2_map[k]
		_, ok1 := result[k]
		// 存在于nums2中，不存在与result中，将加入result
		if ok && !ok1 {
			result[k] = 1
		}
		// 存在于num2中，也存在于result中，不做处理；
		// if ok && !ok {
		// }
	}

	var result_slice []int
	for k := range result {
		result_slice = append(result_slice, k)
	}
	
	return result_slice

}
