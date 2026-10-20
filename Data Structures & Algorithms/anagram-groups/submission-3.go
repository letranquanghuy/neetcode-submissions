func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int

		for i := 0; i < len(s); i++ {
			count[s[i] - 'a']++
		}

		groups[count] = append(groups[count], s)
	}

	results := make([][]string, 0, len(groups))

	for _, g := range groups {
		results = append(results, g)
	}
	return results
}
