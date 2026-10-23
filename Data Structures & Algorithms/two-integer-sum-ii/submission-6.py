class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        for i in range(len(numbers)):
            difference = target - numbers[i]
            j = self.binary_search(numbers, difference, i)
            if j is not None:
                return [i+1, j+1]

    def binary_search(self, numbers: List[int], target: int, not_be: int):
        left = 0
        right = len(numbers) - 1

        while left <= right:
            middle = (left + right) // 2
            if numbers[middle] == target and middle != not_be:
                return middle
            
            if numbers[middle] > target:
                right = middle - 1
            else:
                left = middle + 1
        return None