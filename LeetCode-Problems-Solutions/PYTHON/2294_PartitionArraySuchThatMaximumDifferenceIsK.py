# NAME : Shubham Arora
# DATE : 22/Oct/2022

from typing import List

class Solution:
    def partitionArray(self, nums: List[int], k: int) -> int:
        nums.sort()
        ans = 1
        start = nums[0]
        for i in range(1, len(nums)):
            diff = nums[i] - start
            if diff > k:
                ans += 1
                start = nums[i]
        return ans

# Example usage
solution = Solution()

# Test case 1
nums = [1, 3, 2, 4, 7]
k = 2
print(solution.partitionArray(nums, k))  # Output should be 2

# Test case 2
nums = [1, 5, 3, 2, 8]
k = 3
print(solution.partitionArray(nums, k))  # Output should be 2
