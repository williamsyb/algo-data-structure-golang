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

func InOrder(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}

	var (
		cur   *TreeNode
		stack []*TreeNode
		node  *TreeNode
	)
	cur = root
	stack = []*TreeNode{}

	for len(stack) != 0 || cur != nil {
		// 1.先将左分支都存入stack
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		// 2.将左分支拿出一个存入result
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.val)
		// 3.每次存入后都查一下是否有右分支，
		//有的话设置cur指针，在下一轮循环中再次试图将其所有左分支存入stack
		if node.right != nil {
			cur = node.right
		}
	}
	return
}

func main() {
	for _, val := range InOrder(InitTree()) {
		fmt.Println(val)
	}
}
