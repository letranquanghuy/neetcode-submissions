class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        seen = {}
        begin = longest = 0
        for end, c in enumerate(s):
            if c in seen and seen[c] >= begin:
                begin = seen[c] + 1
            seen[c] = end
            longest = max(longest, end - begin + 1)
        return longest