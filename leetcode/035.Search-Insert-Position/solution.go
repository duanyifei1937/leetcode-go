package leetcode

func searchInsert035(nums []int, target int) int {
	low, high := 0, len(nums)-1
	ans := len(nums)

	for low <= high {
		middle := low + (high-low)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			ans = middle
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return ans

}

// []int{1,3,5,6}, 2)
