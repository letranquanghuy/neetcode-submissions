import "slices"

func hasDuplicate(nums []int) bool {
	n := len(nums)
    unique := make([]int, 0, n)
	for _, value := range nums {
		if slices.Contains(unique, value) {
			return true
		}
		unique = append(unique, value)
	}
	return false
}
