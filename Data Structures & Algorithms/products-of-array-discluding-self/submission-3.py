class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        result = [1] * n

        # Lượt 1: result[i] = tích các phần tử đứng TRƯỚC i
        prefix = 1
        for i in range(n):
            result[i] = prefix      # gán trước, nhân nums[i] vào sau
            prefix *= nums[i]

        # Lượt 2: nhân thêm tích các phần tử đứng SAU i
        suffix = 1
        for i in range(n - 1, -1, -1):
            result[i] *= suffix
            suffix *= nums[i]

        return result