package leetcode

// 翻转string内单词分三步
// 1. 删除多余空格；
// 2. 全部翻转；
// 3. 单词翻转；

func reverseWords(s string) string {
	slow, fast := 0, 0
	ss := []byte(s)

	// 删除头部空格
	for len(ss) > 0 && fast < len(s) && ss[fast] == ' ' {
		fast++
	}

	// 删除中间多于空格：保证间隔只有一个空格
	for ; fast < len(ss); fast++ {
		if fast-1 > 0 && ss[fast-1] == ss[fast] && ss[fast] == ' ' {
			continue
		}
		ss[slow] = ss[fast]
		slow++
	}

	// 删除尾部空格
	if slow-1 > 0 && ss[slow-1] == ' ' {
		ss = ss[:slow-1]
	} else {
		ss = ss[:slow]
	}

	// 全部翻转
	reverse(&ss, 0, len(ss)-1)

	// 单词范围翻转=>正转
	// i -- words start;
	// j -- words end;
	i := 0
	for i < len(ss) {
		j := i
		for ; j < len(ss) && ss[j] != ' '; j++ {
		}
		reverse(&ss, i, j-1)
		i = j
		i++
	}
	return string(ss)
}

func reverse2(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}
