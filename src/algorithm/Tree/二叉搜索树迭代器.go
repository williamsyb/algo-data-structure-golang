package main

/*
173. 二叉搜索树迭代器
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。

调用 next() 将返回二叉搜索树中的下一个最小的数。
        7
       / \
      3   15
         /  \
        9    20
示例：
BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false

提示：
next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	// root *TreeNode
	arr    []int
	curLoc int
	size   int
}

func inOrder(root *TreeNode) (res []int) {
	var (
		stack []*TreeNode
		node  *TreeNode
		cur   *TreeNode
	)
	cur = root
	stack = []*TreeNode{}
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			cur = node.Right
		}
	}
	return
}

func Constructor(root *TreeNode) BSTIterator {
	var (
		arr []int
	)
	arr = inOrder(root)

	return BSTIterator{
		arr:    arr,
		curLoc: 0,
		size:   len(arr),
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	var res int
	res = this.arr[this.curLoc]
	this.curLoc++
	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.curLoc < this.size
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {

}
