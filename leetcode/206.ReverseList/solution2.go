package leetcode


type ListNode3 struct {
    Val int
    Next *ListNode3
}

func reverseList3(head *ListNode3) *ListNode3 {
	var pre *ListNode3
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
