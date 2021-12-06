package leetcode

import "math"

func minSubArrayLen(target int, nums []int) int {
	var sum int
	start, end := 0, 0
	ans := math.MaxInt32
	n := len(nums)

	for end < n {
		sum += nums[end]
		for sum >= target {
			ans = min209(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}

	return ans

}

func min209(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
