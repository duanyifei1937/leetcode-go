package leetcode

type MyQueue struct {
	instack, outstack []int
}

// 两个栈实现一个队列，第一个back、第二个stack;
func Constructor() MyQueue {
	return MyQueue{
		instack:  make([]int, 0),
		outstack: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.instack = append(q.instack, x)
}

// 将instack所有值给outstack;
func (q *MyQueue) in2out() {
	for len(q.instack) > 0 {
		q.outstack = append(q.outstack, q.instack[len(q.instack)-1])
		q.instack = q.instack[:len(q.instack)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outstack) == 0 {
		q.in2out()
	}
	x := q.outstack[len(q.outstack)-1]
	q.outstack = q.outstack[:len(q.outstack)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.outstack) == 0 {
		q.in2out()
	}
	return q.outstack[len(q.outstack)-1]
}

// 判断是否队列为空；
func (q *MyQueue) Empty() bool {
	return len(q.instack) == 0 && len(q.outstack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
