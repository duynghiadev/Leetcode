# Name: Tanya Gupta
# Date: 02/10/2022

from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums_hash = {}
        for i, num in enumerate(nums):
            potentialMatch = target - num
            if potentialMatch in nums_hash:
                return [nums_hash[potentialMatch], i]
            nums_hash[num] = i

# Example usage
if __name__ == "__main__":
    nums = [2, 7, 11, 15]
    target = 9
    solution = Solution()
    result = solution.twoSum(nums, target)
    print(result)  # Output: [0, 1]
