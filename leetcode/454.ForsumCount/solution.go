package leetcode

// 四元组

// 超出时间限制
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var count int
	nums4_map := make(map[int]int)
	for i := range nums4 {
		nums4_map[nums4[i]] += 1
	}

	for i := range nums1 {
		for j := range nums2 {
			for k := range nums3 {
				res := 0 - nums1[i] - nums2[j] - nums3[k]
				if l, ok := nums4_map[res]; ok {
					count += l
				}
			}
		}
	}
	return count
}

func fourSumCount2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	for _, i := range nums1 {
		for _, j := range nums2 {
			// m[i+j] += 1
			m[i+j]++
		}
	}

	var count int
	for _, i := range nums3 {
		for _, j := range nums4 {
			count += m[-i-j]
			// if k, ok := m[0-i-j]; ok {
			// 	count += k
			// }
		}
	}
	return count

}
