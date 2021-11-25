package leetcode

// 有序数组中、连续的数字起始、结尾下标；
// 思路：
// leftindex <= target; rightindex < target (-1);
// 两次二分查找的组合，得到[x, y], 判断是否符合基本条件，否则[-1, -1]

func searchRange034(nums []int, target int) []int {
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}

}

func searchFirstEqualElement(nums []int, target int) int {
	len := len(nums)
	left, right := 0, len-1
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			if middle == 0 || nums[middle-1] != target {
				return middle
			}
			right = middle - 1
		}
	}
	return -1
}

func searchLastEqualElement(nums []int, target int) int {
	len := len(nums)
	left, right := 0, len-1
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			if middle == 0 || nums[middle+1] != target {
				return middle
			}
			left = middle + 1
		}
	}
	return -1
}
