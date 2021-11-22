package leetcode

// 利用 位运算 >>1: 右移一位，a = a / 2;
// 位运算 << 1: 左移一位，a = a * 2;
func search704(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		// mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
