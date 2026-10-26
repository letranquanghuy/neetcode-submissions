class Solution:
    def maxArea(self, heights: List[int]) -> int:
        l = 0
        r = len(heights) - 1
        max_a = 0
        while l < r:
            edge_1 = (r-l)

            if heights[l] < heights[r]:
                edge_2 = heights[l]
                l += 1
            elif heights[l] > heights[r]:
                edge_2 = heights[r]
                r -= 1
            else:
                edge_2 = heights[r]
                r -= 1
                l += 1
            max_a = max(max_a, edge_1*edge_2)
        return max_a