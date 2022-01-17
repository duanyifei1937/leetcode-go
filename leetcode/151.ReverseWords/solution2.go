package leetcode

func reverseWords2(s string) string {
	ss := []byte(s)

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

	// 全部翻转
	reverse3(&ss, 0, len(ss) -1)

	// 单词翻转
	i := 0
	for i < len(ss) {
		j := i
		for ; j < len(ss) && ss[j] != ' '; j++ {
		}
		reverse3(&ss, i, j-1)
		i = j
		i++
	}
	return string(ss)
}

func reverse3(b *[]byte, left, right int) {
	for left < right {
		(*b)[right], (*b)[left] = (*b)[left], (*b)[right]
		left++
		right--
	}
}