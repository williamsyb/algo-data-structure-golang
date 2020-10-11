package main

import "fmt"

/*
105. 从前序与中序遍历序列构造二叉树
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	var (
		rootNum   int
		rootIndex int
		root      *TreeNode
		i         int
		leftArr   []int
		rightArr  []int
	)
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	rootNum = preorder[0]
	for i = 0; i < len(inorder); i++ {
		if rootNum == inorder[i] {
			rootIndex = i
			break
		}
	}
	root = &TreeNode{
		Val: rootNum,
	}
	leftArr = inorder[:rootIndex]
	rightArr = inorder[rootIndex+1:]
	root.Left = buildTree(preorder[1:1+len(leftArr)], leftArr)
	root.Right = buildTree(preorder[1+len(leftArr):], rightArr)
	return root
}

func main() {
	a := []int{1, 3, 5, 7, 4}
	fmt.Println(a[1:1])
}
