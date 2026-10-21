func topKFrequent(nums []int, k int) []int {
	// Count
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	// Bucket sort
	max_frequent := len(nums)
	buckets := make([][]int, max_frequent+1)

	for num, f := range counts {
		buckets[f] = append(buckets[f], num)
	}

	// Choose top K frequent
	result := make([]int, 0, k)

	for f := max_frequent; f >= 1 && len(result) < k; f-- {
		for _, num := range buckets[f] {
			result = append(result, num)

			if len(result) == k {
				break
			}
		}
	}
	return result
}
