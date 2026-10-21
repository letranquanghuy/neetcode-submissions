type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result strings.Builder
	for _, str := range strs {
		// strconv.Itoa đổi int thành string (Integer TO Ascii): 4 -> "4".
		// len(str) là số BYTE của chuỗi, và decode cũng sẽ cắt theo byte,
		// nên hai bên khớp nhau kể cả khi chuỗi có tiếng Việt hay emoji.
		result.WriteString(strconv.Itoa(len(str)))

		result.WriteByte('#')
		result.WriteString(str)

	}
	return result.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		result = append(result, encoded[j+1:j+1+length])
		i = j+1+length
	}
	return result
}
