package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
让快指针一次前进两步，慢指针一次前进一步，当快指针到达链表尽头时，慢指针就处于链表的中间位置
当链表的长度是奇数时，slow 恰巧停在中点位置；如果长度是偶数，slow 最终的位置是中间偏右

重要作用之一：
寻找链表中点的一个重要作用是对链表进行归并排序。
回想数组的归并排序：求中点索引递归地把数组二分，最后合并两个有序数组。
对于链表，合并两个有序链表是很简单的，难点就在于二分。


*/
func findMiddle(head *ListNode) *ListNode {
	var (
		fast *ListNode
		slow *ListNode
	)
	fast = head
	slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
