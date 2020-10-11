package main

import "fmt"

/*
654. 最大二叉树
给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：

二叉树的根是数组中的最大元素。
左子树是通过数组中最大值左边部分构造出的最大二叉树。
右子树是通过数组中最大值右边部分构造出的最大二叉树。
通过给定的数组构建最大二叉树，并且输出这个树的根节点。



示例 ：

输入：[3,2,1,6,0,5]
输出：返回下面这棵树的根节点：

      6
    /   \
   3     5
    \    /
     2  0
       \
        1
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var (
		//num int
		maxIndex int
		node     *TreeNode
		maxNum   int
		i        int
		left     *TreeNode
		right    *TreeNode
	)
	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}
	if len(nums) == 0 {
		return nil
	}
	maxNum = nums[0]
	maxIndex = 0
	for i = 1; i < len(nums); i++ {

		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIndex = i
		}
	}
	node = &TreeNode{
		Val: maxNum,
	}
	left = constructMaximumBinaryTree(nums[:maxIndex])
	right = constructMaximumBinaryTree(nums[maxIndex+1:])
	node.Left = left
	node.Right = right
	return node

}

func main() {
	var root *TreeNode
	root = constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(root)
}
