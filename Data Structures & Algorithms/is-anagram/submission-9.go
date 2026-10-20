func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// Key của map là byte để khớp với kiểu của s[i].
	// value là số lần xuất hiện.
	counts := make(map[byte]int)

	for i := range len(s) {
		counts[s[i]]++
		counts[t[i]]--
	}
	// Không cần so độ dài trước: nếu khác nhau, sẽ có ô khác 0 ở bước dưới.
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}