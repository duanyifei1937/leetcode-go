package leetcode

import "sort"

func threeSum2(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	
	sort.Ints(nums)

	if nums[0] > 0 {
		return res
	}

	// 第一个元素循环
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]

		if i > 0 && n1 == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res

}
