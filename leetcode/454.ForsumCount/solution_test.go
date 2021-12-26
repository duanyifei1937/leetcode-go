package leetcode

import "testing"

func TestFourSumCount(t *testing.T) {
	var nums1, nums2, nums3, nums4 []int
	nums1 = []int{1, 2}
	nums2 = []int{-2, -1}
	nums3 = []int{-1, 2}
	nums4 = []int{0, 2}
	t.Log(fourSumCount(nums1, nums2, nums3, nums4))


	nums1 = []int{0,1,-1}
	nums2 = []int{-1,1,0}
	nums3 = []int{0,0,1}
	nums4 = []int{-1,1,1}

	t.Log(fourSumCount(nums1, nums2, nums3, nums4))
}


func TestFourSumCount2(t *testing.T) {
	var nums1, nums2, nums3, nums4 []int
	nums1 = []int{1, 2}
	nums2 = []int{-2, -1}
	nums3 = []int{-1, 2}
	nums4 = []int{0, 2}
	t.Log(fourSumCount2(nums1, nums2, nums3, nums4))


	nums1 = []int{0,1,-1}
	nums2 = []int{-1,1,0}
	nums3 = []int{0,0,1}
	nums4 = []int{-1,1,1}

	t.Log(fourSumCount2(nums1, nums2, nums3, nums4))
}

