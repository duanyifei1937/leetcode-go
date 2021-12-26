package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]] = i
	}
	for k, v := range m {
		if k1, ok := m[target-k]; ok{
			return []int{v, k1}
		}
	}
	return []int{}
}
