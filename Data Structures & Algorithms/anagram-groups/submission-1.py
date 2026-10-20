class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        seen = {}
        for value in strs:
            sorted_value = str(sorted(value))
            if sorted_value in seen:
                seen[sorted_value].append(value)
                continue
            seen[sorted_value] = [value]
        
        return [v for v in seen.values()]