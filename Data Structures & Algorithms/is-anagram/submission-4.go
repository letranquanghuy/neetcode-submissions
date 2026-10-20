func isAnagram(s string, t string) bool {
	// map[rune]int: key là rune (kiểu số nguyên đại diện 1 ký tự Unicode),
	// value là số lần xuất hiện.
	counts := make(map[rune]int)

	// range trên string cho (vị trí byte, rune). Ta chỉ cần rune.
	for _, r := range s {
		// Key chưa có thì counts[r] trả về 0 (zero value, như ta đã nói
		// ở câu hỏi trước), nên ++ chạy đúng mà không cần kiểm tra ok.
		counts[r]++
	}
	for _, r := range t {
		counts[r]--
	}
	// Không cần so độ dài trước: nếu khác nhau, sẽ có ô khác 0 ở bước dưới.
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}