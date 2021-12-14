package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func comporeListTable(a, b *ListNode) bool {
	return a == b
	
}

func main() {
	var a0 = &ListNode{}
	var a1 = new(ListNode)
	a1.Val = 1
	var a2 = new(ListNode)
	a2.Val = 2
	var a3 = new(ListNode)
	a3.Val = 3
	a0.Next = a1
	a1.Next = a2
	a2.Next = a3

	var b0 = new(ListNode)
	var b1 = new(ListNode)
	b1.Val = 1
	var b2 = new(ListNode)
	b2.Val = 2
	var b3 = new(ListNode)
	b3.Val = 3
	b0.Next = b1
	b1.Next = b2
	b2.Next = b3

	fmt.Println(comporeListTable(a0, b0))
}
