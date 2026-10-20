class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums_dict = {}

        for index, num in enumerate(nums):
            nums_dict.setdefault(num, [])
            nums_dict[num].append(index)

        i, j = 0, 0
        for num in nums:
            if nums_dict.get(target - num) is not None:
                if target - num == num:
                    if len(nums_dict[num]) > 1:
                        i = nums_dict[num][0] 
                        j = nums_dict[num][1] 
                else:
                    i = nums_dict[num][0]      
                    j = nums_dict[target - num][0]      
                    break

        return sorted([i, j])