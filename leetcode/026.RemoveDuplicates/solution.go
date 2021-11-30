package leetcode

func removeDuplicates26(nums []int) int {
	slow := 1
	for fast := 1; fast < len(nums); fast ++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

