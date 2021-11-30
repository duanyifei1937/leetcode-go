package leetcode

func moveZeroes283(nums []int) {
	if len(nums) == 0 {
		return
	}
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
