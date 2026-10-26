class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        res = []
        nums.sort()

        for i in range(0, len(nums)-1):
            l = i + 1
            r = len(nums) - 1

            while l < r:
                s = nums[i] + nums[l] + nums[r]
                data = [nums[i], nums[l], nums[r]]
                if s == 0 and data not in res:
                    res.append(data)
                    l += 1
                elif s < 0:
                    l += 1
                else:
                    r -= 1

        return res