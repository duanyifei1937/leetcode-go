package leetcode

func isPerfectSquare367(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (right + left) >> 1
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false

}
