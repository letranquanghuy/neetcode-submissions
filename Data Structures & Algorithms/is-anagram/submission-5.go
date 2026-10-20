func isAnagram(s string, t string) bool {
	// len(s) trả về số BYTE của chuỗi. Với bài này chỉ có a-z (1 byte/ký tự)
	// nên số byte = số ký tự. Hai chuỗi khác độ dài thì chắc chắn không phải
	// anagram, trả về false ngay để khỏi tốn công đếm.
	if len(s) != len(t) {
		return false
	}

	// s and t consist of lowercase English letters.
	// [26]int là MẢNG có kích thước cố định 26 phần tử kiểu int.
	// Khác slice ([]int): kích thước của mảng nằm trong kiểu và không đổi được.
	// Go tự gán mọi ô bằng zero value (0 với int), nên không cần khởi tạo.
	// Ta dùng 26 ô cho 26 chữ cái: ô 0 = 'a', ô 1 = 'b', ..., ô 25 = 'z'.
	var counts [26]int

	// Vòng for kiểu C: khởi tạo; điều kiện; bước nhảy. Go không có ngoặc tròn.
	// Đây là vòng for "ba phần", dùng khi cần chỉ số i.
	for i := 0; i < len(s); i++ {

		// s[i] lấy ra 1 BYTE tại vị trí i (kiểu byte, không phải string).
		// 'a' là ký tự hằng, giá trị 97. Phép trừ s[i]-'a' đổi chữ cái
		// thành chỉ số 0..25 (ví dụ 'c' - 'a' = 2).
		// Dùng được làm chỉ số mảng vì kết quả là số nguyên.
		//
		// ++ và -- trong Go là CÂU LỆNH, không phải biểu thức
		// (không viết được x = counts[i]++ như C hay Java).
		counts[s[i]-'a']++ // gặp ký tự của s: cộng 1
		counts[t[i]-'a']-- // gặp ký tự của t: trừ 1
	}

	// "for _, c := range counts" duyệt qua mảng: mỗi vòng cho (chỉ số, giá trị).
	// "_" bỏ qua chỉ số vì ta chỉ cần giá trị (Go cấm biến khai báo mà không dùng).
	for _, c := range counts {
		// Nếu là anagram, mỗi ký tự được cộng và trừ số lần bằng nhau
		// nên mọi ô phải bằng 0. Còn ô nào khác 0 là có ký tự lệch số lượng.
		if c != 0 {
			return false
		}
	}

	return true
}