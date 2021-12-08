package main

import (
	"fmt"
)

// 二维数组内存地址并不连续，同行连续
func main() {

	aa := [3][3]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	fmt.Println(&aa[0][0])
	fmt.Println(&aa[0][1])
	fmt.Println(&aa[0][2])
	fmt.Println(&aa[1][0])
	fmt.Println(&aa[1][1])
	fmt.Println(&aa[1][2])
	fmt.Println(&aa[2][0])
	fmt.Println(&aa[2][1])
	fmt.Println(&aa[2][2])

}
