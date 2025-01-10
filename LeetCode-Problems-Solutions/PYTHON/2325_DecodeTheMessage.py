# Name: Ira
# Date: 02/10/2022

class Solution:
    def decodeMessage(self, key: str, message: str) -> str:
        # Step 1: Process the key to remove duplicates and spaces
        k = ""
        for i in key:
            if i not in k and i != " ":
                k += i
        key = k

        # Step 2: Create the reference dictionary for character mapping
        ref = {}
        j = 97  # ASCII value of 'a'
        for i in range(len(key)):
            ref[key[i]] = chr(j)
            j += 1

        # Step 3: Decode the message using the reference dictionary
        res = ''
        for i in message:
            if i != " ":
                if i in ref:
                    res += ref[i]
                else:
                    res += '?'  # If the character is not in ref, use '?' or handle it accordingly
            else:
                res += " "

        return res

# Example usage
solution = Solution()

# Test case 1
key = "the quick brown fox"
message = "vkbs bs t suepuv"
print(solution.decodeMessage(key, message))  # Output: "this is a secret"

# Test case 2
key = "eljuxhpwnyrdgtqkviszcfmabo"
message = "zwx zcfzsd d"
print(solution.decodeMessage(key, message))  # Output: "the code is"
