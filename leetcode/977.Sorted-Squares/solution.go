package leetcode

func sortedSquares(nums []int) []int {
	zeroi := 0
	n := len(nums)

	for i := 0; i < n && nums[i] < 0; i++ {
		zeroi = i
	}

	ans := make([]int, 0, n)

	for i, j := zeroi, zeroi+1; i >= 0 || j < n; {
		if i < 0 {
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == n {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] < nums[j]*nums[j] {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else {
			ans = append(ans, nums[j]*nums[j])
			j++
		}
	}
	return ans
}