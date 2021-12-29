package sword

func replaceSpace(s string) string {
	ss := []byte(s)
	var res []byte
	for i := range ss {
		if s[i] == ' ' {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, ss[i])
		}
	}
	return string(res)
}

func replaceSpace2(s string) string {
	ss := []byte(s)
	length := len(ss)
	spaceCount := 0

	for i := range ss {
		if ss[i] == ' ' {
			spaceCount++
		}
	}

	tmp := make([]byte, spaceCount*2)
	ss = append(ss, tmp...)
	i := length - 1
	j := length + spaceCount*2 - 1
	for i >= 0 {
		if ss[i] != ' ' {
			ss[j] = ss[i]
			i--
			j--
		} else {
			ss[j] = '0'
			ss[j-1] = '2'
			ss[j-2] = '%'
			i--
			j -= 3
		}
	}
	return string(ss)

}
