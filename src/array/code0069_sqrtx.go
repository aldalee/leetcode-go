// https://leetcode.cn/problems/sqrtx/
package main

func mySqrt(x int) int {
	arr := [65536]int{}
	for i := 1; i < len(arr); i++ {
		arr[i] = i * i
	}
	return binarySearchX(arr, x)
}
func binarySearchX(nums [65536]int, target int) int {
	l, r := 0, len(nums)-1
	index := 0
	for l <= r {
		mid := ((r - l) >> 1) + l
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			index = mid
			l = mid + 1
		}
	}
	return index
}
