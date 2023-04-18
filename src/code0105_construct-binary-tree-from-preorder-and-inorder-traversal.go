// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	pre := 0
	in := 0
	max := int(^uint(0) >> 1)
	return helper(preorder, &pre, inorder, &in, max)
}

func helper(preorder []int, pre *int, inorder []int, in *int, max int) *TreeNode {
	if *pre == len(preorder) || *in == len(inorder) || inorder[*in] == max {
		return nil
	}
	root := &TreeNode{Val: preorder[*pre]}
	*pre++
	root.Left = helper(preorder, pre, inorder, in, root.Val)
	*in++
	root.Right = helper(preorder, pre, inorder, in, max)
	return root
}
