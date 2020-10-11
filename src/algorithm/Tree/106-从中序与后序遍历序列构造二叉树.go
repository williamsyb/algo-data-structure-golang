package main

/*
106. 从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	var (
		root      *TreeNode
		rootVal   int
		leftArr   []int
		rightArr  []int
		i         int
		rootIndex int
	)

	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}

	rootVal = postorder[len(postorder)-1]
	root = &TreeNode{
		Val: rootVal,
	}
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}
	leftArr = inorder[:rootIndex]
	rightArr = inorder[rootIndex+1:]
	root.Left = buildTree2(leftArr, postorder[:len(leftArr)])
	root.Right = buildTree2(rightArr, postorder[len(leftArr):len(leftArr)+len(rightArr)])
	return root
}
