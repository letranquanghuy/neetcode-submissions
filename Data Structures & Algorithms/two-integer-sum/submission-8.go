func twoSum(nums []int, target int) []int {
	// map[int]int: key là giá trị số, value là chỉ số của nó trong nums.
	// (Bài trước value là struct{} vì chỉ cần biết "có/không". Ở đây ta
	// cần nhớ cả VỊ TRÍ, nên value là int.)
	// len(nums) là gợi ý dung lượng ban đầu, cấp phát trước một lần cho nhanh.
	seen := make(map[int]int, len(nums))

	// "for i, num := range nums": mỗi vòng cho (chỉ số, giá trị).
	// Lần này ta dùng cả hai nên không cần "_" để bỏ qua.
	for i, num := range nums {

		// Số cần tìm để cộng với num ra target.
		// ":=" khai báo biến mới, Go tự suy ra kiểu int.
		complement := target - num

		// Tra map trả về 2 giá trị: (value, ok).
		// ok = true nếu key complement CÓ trong map.
		// Ở đây ta dùng cả j (chỉ số) lẫn ok, nên phải kiểm tra ok:
		// nếu chỉ nhìn j thì không phân biệt được "không có key" (j = 0, zero value)
		// với "có key và chỉ số thật sự là 0".
		// "if j, ok := ...; ok" là if với câu lệnh khởi tạo: j và ok chỉ
		// sống trong khối if này.
		if j, ok := seen[complement]; ok {
			// []int{j, i} tạo một slice 2 phần tử (kiểu + {} = tạo giá trị,
			// giống struct{}{} lúc trước). j được lưu từ vòng trước nên
			// j < i, tức chỉ số nhỏ đứng trước đúng như đề yêu cầu.
			return []int{j, i}
		}

		// Chưa thấy số bù thì lưu num vào map để các vòng sau còn tìm thấy.
		// Lưu SAU khi kiểm tra, nên một số không bao giờ ghép với chính nó.
		seen[num] = i
	}

	// Go BẮT BUỘC hàm có kiểu trả về phải kết thúc bằng return,
	// nếu không sẽ lỗi biên dịch "missing return", dù ta biết đề đảm bảo
	// luôn có đáp án và dòng này sẽ không bao giờ chạy tới.
	// nil là giá trị "rỗng" của slice (giống None của Python, nhưng chỉ dùng
	// cho slice, map, con trỏ, interface...).
	return nil
}