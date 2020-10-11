package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/*
	  5
    /   \
   3      6
  / \    /
 2   4  7

*/
func InitTree4() *TreeNode {
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

func levelOrder(root *TreeNode) (res []int) {
	var (
		queue []*TreeNode
		node  *TreeNode
	)
	res = []int{}
	queue = []*TreeNode{root}
	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		res = append(res, node.val)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return
}

func main() {
	fmt.Println(levelOrder(InitTree4()))
}
