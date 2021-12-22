package leetcode

func intersection2(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, i := range nums1 {
		m[i] = 1
	}

	var res []int
	for _, v := range nums2 {
		if n, ok := m[v]; ok && n > 0 {
			res = append(res, v)
			m[v]--
		}
	}

	return res
}

func intersection3(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	res := make([]int, 0)

	for _, v := range nums1 {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}

	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
			delete(m, v)
		}
	}
	return res
}
