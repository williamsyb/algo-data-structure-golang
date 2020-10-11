package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func cycleStartLoc(head *ListNode) *ListNode {
	var (
		fast *ListNode
		slow *ListNode
	)
	fast = head
	slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	//第一次相遇时，假设慢指针 slow 走了 k 步，那么快指针 fast 一定走了 2k 步，也就是说比 slow 多走了 k 步（也就是环的长度）
	//设相遇点距环的起点的距离为 m，那么环的起点距头结点 head 的距离为 k - m，也就是说如果从 head 前进 k - m 步就能到达环起点。

	//巧的是，如果从相遇点继续前进 k - m 步，也恰好到达环起点。
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow

}
