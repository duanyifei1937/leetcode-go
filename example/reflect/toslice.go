package main

import (
	"fmt"
	"reflect"
)

func main() {
	data := []string{"one", "two", "three"}
	fmt.Println(test(data))

}

func test(t interface{}) (l []string) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			l = append(l, s.Index(i).String())
		}
	}
	return l
}
