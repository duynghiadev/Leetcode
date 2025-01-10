# Name: Himanshi Garg
# Date: 06/10/2022

class Solution:
    def pushDominoes(self, dominoes: str) -> str:
        prevString = ""
        current = dominoes

        if len(dominoes) == 1:  # If there is only one domino
            return dominoes

        while prevString != current:  # Keep simulating until no changes
            prevString = current
            current = list(prevString)  # Convert string to list to modify it

            for i in range(len(dominoes)):
                if prevString[i] == ".":
                    # If current domino is "." and adjacent ones have forces applied
                    if i > 0 and prevString[i - 1] == "R" and (i == len(dominoes) - 1 or prevString[i + 1] != "L"):
                        current[i] = "R"  # Apply force towards right
                    elif i < len(dominoes) - 1 and prevString[i + 1] == "L" and prevString[i - 1] != "R":
                        current[i] = "L"  # Apply force towards left
                # If the current domino is either "R" or "L", no change needed

            current = "".join(current)  # Convert the list back to a string

        return current

# Example Test Case
s = Solution()
dominoes = "RR.L"
result = s.pushDominoes(dominoes)
print(result)  # Expected Output: "RR.L"