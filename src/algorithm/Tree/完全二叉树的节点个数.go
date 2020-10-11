package main

/*
222. 完全二叉树的节点个数
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	// if root==nil{
	//     return 0
	// }
	// return 1+countNodes(root.Left)+countNodes(root.Right)
	// 用递归的方式写性能会差很多
	var (
		stack []*TreeNode
		node  *TreeNode
		res   int
	)

	stack = []*TreeNode{root}
	for len(stack) != 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res++
		} else {
			continue
		}

		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	return res
}
