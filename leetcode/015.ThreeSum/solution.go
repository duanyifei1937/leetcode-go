package leetcode

func threeSum(nums []int) [][]int {
	m := make(map[int]int)
	res := make([][]int, 0)

	if len(nums) < 3 {
		return [][]int{}
	}

	for i := range nums {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if k, ok := m[-nums[i]-nums[j]]; ok && k > j {
				res = append(res, []int{nums[i], nums[j], nums[k]})

			}
		}
	}
	return res
}

func sortThree(res []int) []int {
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return res
}
