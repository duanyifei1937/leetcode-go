package linktable

type LinkNode struct {
	val  int
	Next *LinkNode
}

func removeElement(head *LinkNode, val int) *LinkNode {
	dummyHead := &LinkNode{}
	dummyHead.Next = head

	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
