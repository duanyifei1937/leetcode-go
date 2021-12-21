package leetcode

func isAnagram242(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1, t1 := make(map[rune]int), make(map[rune]int)

	for _, i := range s {
		s1[i] += 1
	}

	for _, i := range t {
		t1[i] += 1
	}

	for i := range s1 {
		if s1[i] != t1[i] {
			return false
		}
	}
	return true

}
