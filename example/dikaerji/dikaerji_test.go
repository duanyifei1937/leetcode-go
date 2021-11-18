package main

import (
	"testing"
)

func TestErDiKaEr(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"d", "e", "f"}

	c := make([][]string, 0)

	for _, i := range a {
		for _, j := range b {
			c = append(c, []string{i, j})
		}
	}

	for _, i := range c {
		t.Log("--")
		t.Log(i)
		t.Logf("%T", i)
	}

}

func TestThreeDiKaEr(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"d", "e", "f"}
	c := []string{"m", "n"}

	cl := make([][]string, 0)

	for _, i := range a {
		for _, j := range b {
			for _, k := range c {
				cl = append(cl, []string{i, j, k})
			}
		}
	}
	t.Log(cl)
}

func TestDemo(t *testing.T) {
	result := make([][]int, 0)
	result = append(result, []int{1, 2})
	result = append(result, []int{2, 3})

	t.Log(result)
}
