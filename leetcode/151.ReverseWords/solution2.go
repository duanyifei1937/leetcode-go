package leetcode

func reverseWords2(s string) string {
	ss := []byte(s)
	
	// //第一次翻转
	// reverString(ss)

	// // 第二次按空格翻转
	// var start, end int
	// for i :=0; i<len(ss); i++{
	// 	if ss[i] == ' ' {
	// 		end = i
	// 		reverString(ss[start:end])
	// 		start = i + 1
	// 	}
	// }
	// return string(ss)

	slow, fast := 0, 0
	// 删除头部空格
	for len(ss) > 0 && fast < len(ss) && ss[fast] == ' ' {
		fast++
	}

	// 删除多于中间空格
	for ; fast < len(ss); fast++ {
		if fast-1>0 && ss[fast-1] == ss[fast] && ss[fast] == ' ' {
			continue
		}
		ss[slow] = ss[fast]
		slow++
	}

	// 删除尾部多于空格
	if slow -1 > 0 && ss[slow-1] == ' ' {
		ss = ss[:slow-1]
	} else {
		ss = ss[:slow]
	}
	return string(ss)


}



func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[right], (*b)[left] = (*b)[left], (*b)[right]
		left++
		right--
	}
}