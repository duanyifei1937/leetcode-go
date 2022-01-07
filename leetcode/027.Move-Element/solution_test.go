package leetcode

import (
	"fmt"
	"testing"
)

// 	输入：nums = [3,2,2,3], val = 3
// 输出：2, nums = [2,2]

// 输入：nums = [0,1,2,2,3,0,4,2], val = 2
// 输出：5, nums = [0,1,4,0,3]

type question27 struct {
	para27
	ans27
}

type para27 struct {
	one []int
	two int
}

type ans27 struct {
	one int
}

func Test_Problem27(t *testing.T) {
	qs := []question27{
		{
			para27{[]int{1, 0, 1}, 1},
			ans27{1},
		},
		{
			para27{[]int{0, 1, 0, 3, 0, 12}, 0},
			ans27{3},
		},
	}

	for _, q := range qs {
		p := q.para27
		fmt.Printf("[intput]: %v, %v	[output]: %v	[expect]: %v\n", p.one, p.two, removeElement027_2(p.one, p.two), q.ans27.one)
		fmt.Printf("[intput]: %v, %v	[output]: %v	[expect]: %v\n", p.one, p.two, removeElement027(p.one, p.two), q.ans27.one)
	}
}

func TestRemoveElement2(t *testing.T) {
	t.Log(removeElement2([]int{3, 2, 2, 3}, 3))
	t.Log(removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	t.Log(removeElement2([]int{1}, 1))

}
