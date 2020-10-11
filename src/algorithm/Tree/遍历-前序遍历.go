package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func InitTree() *TreeNode {
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

func Solution(root *TreeNode) (result []*TreeNode) {
	var (
		stack []*TreeNode
		node  *TreeNode
	)
	stack = []*TreeNode{root}
	for len(stack) != 0 {

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		stack = append(stack, node.right)
		stack = append(stack, node.left)

		result = append(result, node)
	}
	return
}

func main() {
	for _, item := range Solution(InitTree()) {
		fmt.Println(item.val)
	}
}
