package main

import "fmt"

func search7041(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		middle := low + (high-low)/2
		if nums[middle] == target {
			return middle
		} else if target > nums[middle] {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}

func search7042(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high {
		middle := low + (high-low)/2
		if nums[middle] == target {
			return middle
		} else if target > nums[middle] {
			low = middle + 1
		} else {
			high = middle
		}
	}
	return -1
}

func main() {
	fmt.Println(search7041([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(search7042([]int{-1, 0, 3, 5, 9, 12}, 2))
}
