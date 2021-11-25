package leetcode

func mySqrt069(x int) int {
	left, right := 0, x
	for left <= right {
		middle := (right + left) >> 1
		sqrt := middle * middle
		if sqrt == x || (sqrt < x && x < (middle+1)*(middle+1)) {
			return middle
		} else if sqrt > x {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return 0

}
