class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        parent = {}
        size = {}

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]  # path compression
                x = parent[x]
            return x

        def union(a, b):
            ra, rb = find(a), find(b)
            if ra == rb:
                return
            if size[ra] < size[rb]:
                ra, rb = rb, ra
            parent[rb] = ra
            size[ra] += size[rb]

        for n in set(nums):
            parent[n] = n
            size[n] = 1

        for n in parent:
            if n + 1 in parent:
                union(n, n + 1)

        return max(size.values()) if size else 0