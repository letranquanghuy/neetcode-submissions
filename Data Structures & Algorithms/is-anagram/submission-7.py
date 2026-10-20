class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        chars = dict()

        for i in range(len(s)):
            chars.setdefault(s[i], 0)
            chars.setdefault(t[i], 0)

            chars[s[i]] += 1
            chars[t[i]] -= 1
        
        result = not any(v > 0 for v in chars.values())
        return result