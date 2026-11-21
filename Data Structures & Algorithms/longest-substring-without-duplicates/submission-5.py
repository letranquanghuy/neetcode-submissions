class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        
        begin = 0
        end = 0
        longest = 0
        seen = {}
        while end < len(s):
            if s[end] in seen and seen[s[end]] >= begin and seen[s[end]] < end:
                longest = max(longest, end - begin)
                begin = seen[s[end]] + 1
            seen[s[end]] = end
            end += 1
        longest = max(longest, end - begin)
        return longest
