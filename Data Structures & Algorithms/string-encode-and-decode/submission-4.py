class Solution:

    def encode(self, strs: List[str]) -> str:
        # idea treat each string as a dataframe which 1st part is metadata include length of string
        # 2nd part is this string

        result = ""
        for s in strs:
            result = result + str(len(s)) + "#" + s
        return result
    
    def decode(self, s: str) -> List[str]:
        result = []
        i = 0
        while i < len(s):
            j = i
            while s[j] != "#":
                j += 1
            
            length = int(s[i:j])
            result.append(s[j+1: j+1+length])
            i = j+1+length
        return result