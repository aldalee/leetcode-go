// https://leetcode.cn/problems/check-permutation-lcci/
package main

import "sort"

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	b1, b2 := []byte(s1), []byte(s2)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	return string(b1) == string(b2)
}
func main() {

}
