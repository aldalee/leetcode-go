// https://leetcode.cn/problems/assign-cookies/
package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	// 从小到大
	// s 饼干 下标 j
	// g 孩子 下标 i
	i := 0
	// 遍历饼干，看能否满足孩子
	for j := 0; j < len(s) && i < len(g); j++ {
		//如果饼干的大小≥孩子的胃口则给与，否则不给予，继续寻找选一个饼干是否符合
		if s[j] >= g[i] {
			i++
		}
	}
	return i
}
func main() {

}
