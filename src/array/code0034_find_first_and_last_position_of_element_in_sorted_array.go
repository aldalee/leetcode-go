// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
package main

import "math"

func searchRange(nums []int, target int) []int {
	l, r, mid := binarySearch(nums, target)
	if mid == -1 {
		return []int{-1, -1}
	}
	start := binarySearchLeft(nums, l, mid, target)
	end := binarySearchRight(nums, mid, r, target)
	return []int{start, end}
}

func binarySearch(nums []int, target int) (l int, r int, index int) {
	l, r = 0, len(nums)-1
	for l <= r {
		mid := ((r - l) >> 1) + l
		if nums[mid] == target {
			return l, r, mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l, r, -1
}

func binarySearchLeft(nums []int, l int, r int, target int) int {
	index := int(math.Max(float64(l), float64(r)))
	for l <= r {
		mid := ((r - l) >> 1) + l
		if nums[mid] == target {
			index = mid
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	return index
}

func binarySearchRight(nums []int, l int, r int, target int) int {
	index := int(math.Min(float64(l), float64(r)))
	for l <= r {
		mid := ((r - l) >> 1) + l
		if nums[mid] == target {
			index = mid
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	return index
}
