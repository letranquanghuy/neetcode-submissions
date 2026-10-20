func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	n := len(s)
	s_chars := make(map[rune]int, n)
	t_chars := make(map[rune]int, n)
	s_rune := []rune(s)
	t_rune := []rune(t)

	for i := 0; i < n; i++ {
		s_chars[s_rune[i]]++
		t_chars[t_rune[i]]++
	}
	for k, v := range s_chars {
		if w, ok := t_chars[k]; !ok || v != w {
			return false
		}
	}

	return true
}
