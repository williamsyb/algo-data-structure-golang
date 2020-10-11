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

func PostOrder(root *TreeNode) (res []*TreeNode) {
	var (
		stack  []*TreeNode
		node   *TreeNode
		result []*TreeNode
		i      int
	)
	stack = []*TreeNode{root}
	for {
		if len(stack) == 0 {
			break
		}
		node = stack[len(stack)-1]

		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		result = append(result, node)
		stack = append(stack, node.left)
		stack = append(stack, node.right)

	}
	for i = len(result) - 1; i >= 0; i-- {
		res = append(res, result[i])
	}
	return
}

func main() {
	for _, node := range PostOrder(InitTree()) {
		fmt.Println(node.val)
	}
}
