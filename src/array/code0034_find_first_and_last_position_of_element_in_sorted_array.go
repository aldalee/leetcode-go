// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
package main

import "sort"

func searchRange(nums []int, target int) []int {
	L := sort.SearchInts(nums, target)
	if L == len(nums) || nums[L] != target {
		return []int{-1, -1}
	}
	R := sort.SearchInts(nums, target+1) - 1
	return []int{L, R}
}
