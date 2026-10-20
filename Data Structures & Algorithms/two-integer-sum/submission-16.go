func twoSum(nums []int, target int) []int {
	nums_map := make(map[int][]int)

	for index, num := range nums {
		nums_map[num] = append(nums_map[num], index)
	}
	i := 0
	j := 0

	for _, num := range nums {
		if len(nums_map[target - num]) != 0 {
			if (target - num) == num {
				if len(nums_map[target - num]) > 1 {
					i = nums_map[num][0]
					j = nums_map[num][1]
					break
				}
			} else {
				i = nums_map[num][0]
				j = nums_map[target - num][0]
				break
			}
		}
	}
	if i > j {
		i, j = j, i
	}
	return []int{i, j}

}
