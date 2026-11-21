class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        seen = {}
        begin = longest = 0
        for end, c in enumerate(s):
            if c in seen and seen[c] >= begin:
                begin = seen[c] + 1
            if (current := end - begin + 1 ) > longest:
                longest = current
            seen[c] = end
        return longest