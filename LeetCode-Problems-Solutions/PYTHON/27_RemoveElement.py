# Name - Akruti Sarangi
# Date - 20/10/2022
# According to the problem we have to remove element from the list and return the count of elements after removing that element

class Solution:
    """
    :type nums: List[int]
    :type val: int
    :rtype: int
    """
    def removeElement(self, nums, val):
        ls = len(nums)
        if ls == 0:
            return ls
        count = 0
        index = 0
        # Count the number of elements that need to be removed
        while index < ls - count:
            if nums[index] == val:
                nums[index] = nums[ls - 1 - count]
                count += 1
            else:
                index += 1
        # Return the number of elements remaining
        return ls - count

if __name__ == "__main__":
    # Begin testing
    s = Solution()
    print(s.removeElement([1], 1))  # Output: 0
    print(s.removeElement([3, 2, 2, 3], 3))  # Output: 2
    print(s.removeElement([0, 1, 2, 2, 3, 0, 4, 2], 2))  # Output: 5
