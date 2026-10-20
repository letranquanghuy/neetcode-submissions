func bb_sort(str string) string {
	chars := []byte(str)

	for i := 0; i < len(chars) - 1; i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] > chars[j] {
				chars[i], chars[j] = chars[j], chars[i]
			}
		}
	}
	return string(chars)
}

func groupAnagrams(strs []string) [][]string {
	seen := make(map[string][]string)

	for _, value := range strs {
		sorted_value := bb_sort(value)
		if _, ok := seen[sorted_value]; ok {
			seen[sorted_value] = append(seen[sorted_value], value)
			continue
		}
		seen[sorted_value] = []string{value}
	}

	result := [][]string{}
	for _, value := range seen {
		result = append(result, value)
	}
	return result
}
