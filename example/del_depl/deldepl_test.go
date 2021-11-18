// 去重

package main

import (
	"testing"
)

func TestDelDupl(t *testing.T) {
	a := []string{"a", "b", "c", "a"}
	t.Log(a)
}

func TestMapDelDuplicate(t *testing.T) {
	s := make(map[string]string)
	a := []string{"a", "b", "c", "a", "d"}

	for _, i := range a {
		s[i] = ""
	}

	l := make([]string, 0)
	for i := range s {
		l = append(l, i)
	}
	
	t.Log(l)

}






