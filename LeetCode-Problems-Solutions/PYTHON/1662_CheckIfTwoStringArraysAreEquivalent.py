# Name : Simran Kukreja
# Date : 26th Oct 2022

class Solution(object):
    def arrayStringsAreEqual(self, word1, word2):
        """
        :type word1: List[str]
        :type word2: List[str]
        :rtype: bool
        """
        # Join the lists into single strings and compare
        return ''.join(word1) == ''.join(word2)

# Test case
word1 = ["abc", "def"]
word2 = ["abc", "def"]
word3 = ["ab", "cdef"]

solution = Solution()

# Test with word1 and word2
print(solution.arrayStringsAreEqual(word1, word2))  # Should print True

# Test with word1 and word3
print(solution.arrayStringsAreEqual(word1, word3))  # Should print False
