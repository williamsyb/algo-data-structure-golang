package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func InitTree3() *TreeNode {
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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var (
		temp *TreeNode
	)
	temp = root.left
	root.left = root.right
	root.right = temp
	invertTree(root.left)
	invertTree(root.right)
	return root
}

func main() {
	root := InitTree3()
	root = invertTree(root)

}
