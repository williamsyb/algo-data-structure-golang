package main

import "fmt"

/*
一、题目
给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

二、案例
输入:
	  5
    /   \
   3      6
  / \    /
 2   4  7

Target = 9

输出: True（因为存在 2 + 7 = 9）

三、思路
使用中序遍历得到有序数组之后，再利用双指针对数组进行查找。
应该注意到，这一题不能用分别在左右子树两部分来处理这种思想，因为两个待求的节点可能分别在左右子树中。
*/

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func InitTree2() *TreeNode {
	var root *TreeNode
	root = &TreeNode{
		val: 5,
	}
	root.left = &TreeNode{val: 3}
	root.right = &TreeNode{val: 6}
	root.left.left = &TreeNode{val: 2}
	root.left.right = &TreeNode{val: 4}
	root.right.left = &TreeNode{val: 7}
	return root
}

func InOrder2(root *TreeNode) (res []int) {
	var (
		stack []*TreeNode
		cur   *TreeNode
		node  *TreeNode
	)
	stack = []*TreeNode{}
	cur = root
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		node = stack[len(stack)-1]
		res = append(res, node.val)
		stack = stack[:(len(stack) - 1)]
		cur = node.right

	}
	return
}

func Solution2(root *TreeNode, target int) bool {

	var (
		arr []int
		i   int
		j   int
	)
	arr = InOrder2(root)
	fmt.Println("root -> arr:", arr)
	i = 0
	j = len(arr) - 1
	for i < j {
		if arr[i]+arr[j] == target {
			return true
		} else if arr[i]+arr[j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

func main() {
	root := InitTree2()
	fmt.Println(Solution2(root, 13))
}
