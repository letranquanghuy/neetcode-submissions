class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counts = dict()
        for n in nums:
            counts[n] = 1 + counts.get(n, 0)

        max_frequent = len(nums)
        buckets = [[] for _ in range(max_frequent+1)]
        for num, frequent in counts.items():
            buckets[frequent].append(num)       
        
        result = []
        frequent = max_frequent
        while frequent >= 1 and len(result) < k:
            for num in buckets[frequent]:
                if len(result) < k:
                    result.append(num)
            frequent -= 1
        
        return result