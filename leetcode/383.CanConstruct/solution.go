package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, i := range ransomNote {
		m1[i]++
	}

	for _, i := range magazine {
		m2[i]++
	}

	for k, v := range m1 {
		if v1, ok := m2[k]; !ok || v > v1 {
			return false
		}
	}

	return true

}

func canConstruct2(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, v := range magazine {
		record[v-'a']++
	}
	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}
	return true
}
