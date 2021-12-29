package leetcode

func reverseWords(s string) string {
	ss := []byte(s)
	var fast int

	for slow := 0; slow < len(ss); {
		if ss[slow] == ' ' {
			for fast = slow + 1; fast < len(ss); fast++ {
				if ss[fast] != ' ' {
					break
				}
			}
			ss[slow+1] = ss[fast]
			slow += 2
		} else {
			slow++
		}
	}
	return string(ss)

}
