class Solution:
    def isPalindrome(self, s: str) -> bool:
        clean_s = ""
        for char in s:
            if char.isalnum():
                clean_s += char.lower()
       
        middle = len(clean_s) // 2
        if len(clean_s) % 2 == 0:
            return clean_s[:middle] == clean_s[len(clean_s)-1: middle-1: -1]
        else:
            return clean_s[:middle] == clean_s[len(clean_s)-1: middle: -1]