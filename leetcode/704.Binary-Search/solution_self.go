package leetcode

//测试不全面
func search704Self(nums []int, target int) int {
	var count int
	for {
		if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
			return -1
		}

		middle := nums[len(nums)/2]
		if middle == target {
			return len(nums)/2 + count
		} else if target > middle {
			count += len(nums) / 2
			nums = nums[len(nums)/2:]
		} else {
			count += len(nums) / 2
			nums = nums[:len(nums)/2]
		}
	}
}
