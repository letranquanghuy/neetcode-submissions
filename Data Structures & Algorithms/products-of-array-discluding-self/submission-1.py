class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        result = []
        prefix_product = [1]*len(nums)
        suffix_product = [1]*len(nums)

        for i in range(len(nums)):
            reverse_i = -(i+1)
            if i == 0:
                prefix_product[i] = nums[i]
                suffix_product[reverse_i] = nums[reverse_i]
                continue
            prefix_product[i] = prefix_product[i-1]*nums[i]
            suffix_product[reverse_i] = suffix_product[reverse_i+1]*nums[reverse_i]

        for i in range(len(nums)):
            if i == 0:
                result.append(suffix_product[i+1])
            elif i == len(nums)-1:
                result.append(prefix_product[i-1])
            else:
                result.append(prefix_product[i-1]*suffix_product[i+1])
        return result