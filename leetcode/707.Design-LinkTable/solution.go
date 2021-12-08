package leetcode

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

type MyLinkedList struct {
	//虚拟头结点定义
	dummy *Node
}

// 哑节点，头、尾都是自己；
// 初始化链表, 只有一个哑节点
func Constructor() MyLinkedList {
	rear := &Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	rear.Next = rear
	rear.Pre = rear
	return MyLinkedList{rear}
}

// 获取链表中第 index 个节点的值。如果索引无效，则返回-1
func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	// 存在后续元素
	for head != this.dummy && index > 0 {
		index--
		head = head.Next
	}
	// 不存在元素
	if 0 != index {
		return -1
	}
	return head.Val
}

// 在链表的第一个元素之前添加一个值为 val 的节点
func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Next: dummy.Next,
		Pre:  dummy,
	}
	dummy.Next.Pre = node
	dummy.Next = node
}

// 将值为 val 的节点追加到链表的最后一个元素
func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	rear := &Node{
		Val:  val,
		Next: dummy,
		Pre:  dummy.Pre,
	}

	// 两步不能反
	dummy.Pre.Next = rear
	dummy.Pre = rear
}

// 在链表中的第 index 个节点之前添加值为 val  的节点
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	head := this.dummy.Next
	for head != this.dummy && index > 0 {
		head = head.Next
		index--
	}

	node := &Node{
		Val:  val,
		Next: head,
		Pre:  head.Pre,
	}

	head.Pre.Next = node
	head.Pre = node
}

// 删除指定index的元素
func (this *MyLinkedList) DeleteAtIndex(index int) {
	//链表为空
	if this.dummy.Next == this.dummy {
		return
	}

	head := this.dummy.Next
	for head.Next != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	if index == 0 {
		head.Next.Pre = head.Pre
		head.Pre.Next = head.Next
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
