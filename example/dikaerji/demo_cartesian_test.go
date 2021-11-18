package main

import (
	"testing"

	"github.com/schwarmco/go-cartesian-product"
)

func TestCartesian(t *testing.T) {
	// a := []string{"a", "b", "c"}
	// b := []string{"d", "e", "f"}

	a1 := []interface{}{"a", "b", "c"}
	b1 := []interface{}{"d", "e", "f"}

	r := cartesian.Iter(a1, b1)

	for i := range r {
		t.Logf("%T", i)
		t.Log(i...)
		t.Log(i[0], i[1])
		break
	}
}
