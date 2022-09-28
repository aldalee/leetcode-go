// https://leetcode.cn/problems/search-insert-position/
package main

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}
	length := len(nums)
	if target > nums[length-1] {
		return length
	}
	l, r := 0, length-1
	index := -1
	for l <= r {
		mid := ((r - l) >> 1) + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			index = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return index
}
