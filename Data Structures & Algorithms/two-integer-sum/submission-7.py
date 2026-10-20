class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = {}  # giá trị -> chỉ số
        for i, num in enumerate(nums):
            if target - num in seen:
                return [seen[target - num], i]
            seen[num] = i