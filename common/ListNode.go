package common

type ListNode struct {
	Val  int
	Next *ListNode
}

// Nums2List convert []int to List
func Nums2List(nums []int) *ListNode {
	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

func Nums2ListWithCycle(nums []int, pos int) *ListNode {
	head := Nums2List(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
