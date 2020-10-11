package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//假设k不会超过链表长度
func lastKthNode(head *ListNode, k int) *ListNode {
	/*
		使用快慢指针，让快指针先走 k 步，然后快慢指针开始同速前进。
		这样当快指针走到链表末尾 null 时，慢指针所在的位置就是倒数第 k 个链表节点
		（为了简化，假设 k 不会超过链表长度）
	*/
	var (
		fast *ListNode
		slow *ListNode
	)
	fast = head
	slow = head
	for k > 0 {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
