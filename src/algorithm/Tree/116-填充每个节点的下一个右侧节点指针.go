package main

/*
116. 填充每个节点的下一个右侧节点指针
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



示例：
		1
       / \
      2   3
     / \  / \
    4   5 6  7

         1 -> nil
       /   \
      2  ->  3 -> nil
    /  \    /  \
   4 -> 5-> 6 -> 7 ->nil

输入：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":null,"right":null,"val":4},"next":null,"right":{"$id":"4","left":null,"next":null,"right":null,"val":5},"val":2},"next":null,"right":{"$id":"5","left":{"$id":"6","left":null,"next":null,"right":null,"val":6},"next":null,"right":{"$id":"7","left":null,"next":null,"right":null,"val":7},"val":3},"val":1}

输出：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":{"$id":"4","left":null,"next":{"$id":"5","left":null,"next":{"$id":"6","left":null,"next":null,"right":null,"val":7},"right":null,"val":6},"right":null,"val":5},"right":null,"val":4},"next":{"$id":"7","left":{"$ref":"5"},"next":null,"right":{"$ref":"6"},"val":3},"right":{"$ref":"4"},"val":2},"next":null,"right":{"$ref":"7"},"val":1}

解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。


提示：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type NodeWrapper struct {
	node  *Node
	level int
}

func connect(root *Node) *Node {
	var (
		//levelArr []int
		nodeArr   []*NodeWrapper
		queue     []*NodeWrapper
		node      *NodeWrapper
		lastLevel int
	)
	if root == nil {
		return root
	}
	node = &NodeWrapper{
		node:  root,
		level: 1,
	}
	queue = append(queue, node)
	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		if node != nil {
			nodeArr = append(nodeArr, node)
		} else {
			continue
		}
		if node.node.Left != nil {

			queue = append(queue, &NodeWrapper{
				node:  node.node.Left,
				level: node.level + 1,
			})
		}
		if node.node.Right != nil {
			queue = append(queue, &NodeWrapper{
				node:  node.node.Right,
				level: node.level + 1,
			})
		}

	}
	nodeArr = append(nodeArr, &NodeWrapper{nil, 0})
	lastLevel = 0
	for i := 0; i < len(nodeArr); i++ {
		node = nodeArr[i]
		if node.level != lastLevel {
			lastLevel = node.level
		} else {
			nodeArr[i-1].node.Next = node.node

		}

	}
	return root
}
