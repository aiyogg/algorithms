package main

import "algorithms/common"

/**
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

链接：https://leetcode-cn.com/problems/linked-list-cycle

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
*/

type ListNode = common.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fastHead := head.Next // 快指针
	for fastHead != nil && head != nil && fastHead.Next != nil {
		if fastHead == head { // 快慢指针相遇说明有环
			return true
		}
		// 快指针每次走两步，慢指针每次走一步
		fastHead = fastHead.Next.Next
		head = head.Next
	}
	return false
}
