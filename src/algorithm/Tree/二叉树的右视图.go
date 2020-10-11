package main

import "fmt"

/*
199. 二叉树的右视图
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

		   1            <---
		 /   \
		2     3         <---
		 \     \
		  5     4       <---
         /
        3               <---

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InitTreeNode() *TreeNode {
	var root *TreeNode
	root = &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}
	return root
}

type TreeNodeWrapper struct {
	node  *TreeNode
	Left  *TreeNodeWrapper
	Right *TreeNodeWrapper
	level int
}

func (tn *TreeNodeWrapper) isNull() bool {
	return tn.node == nil
}

func rightSideView(root *TreeNode) []int {
	var (
		rootWrapper *TreeNodeWrapper
		queue       []*TreeNodeWrapper
		node        *TreeNodeWrapper
		levelArr    []*TreeNodeWrapper
		res         []int
		lastLevel   int
	)

	rootWrapper = &TreeNodeWrapper{
		node:  root,
		level: 1,
	}
	if rootWrapper.isNull() {
		return []int{}
	}
	queue = []*TreeNodeWrapper{rootWrapper}
	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		levelArr = append(levelArr, node)
		node.Left = &TreeNodeWrapper{
			node:  node.node.Left,
			level: node.level + 1,
		}
		if !node.Left.isNull() {
			queue = append(queue, node.Left)
		}
		node.Right = &TreeNodeWrapper{
			node:  node.node.Right,
			level: node.level + 1,
		}
		if !node.Right.isNull() {
			queue = append(queue, node.Right)
		}

	}
	lastLevel = 1
	levelArr = append(levelArr, &TreeNodeWrapper{nil, nil, nil, 0})
	for i := 0; i < len(levelArr); i++ {
		node = levelArr[i]
		if node.level != lastLevel {
			res = append(res, levelArr[i-1].node.Val)
			lastLevel++
		}
	}
	return res
}

func main() {
	root := InitTreeNode()
	fmt.Println(rightSideView(root))
}
