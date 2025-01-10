# Name: Aditya Chache
# Date: 06/10/2022


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Edge case: if there's only one node, return None
        if not head.next:
            return None

        # Create a dummy node to handle edge cases (like deleting the head)
        dummy = ListNode()
        dummy.next = head

        # Initialize two pointers: slow and fast
        slow = dummy
        fast = head

        # Move slow by one and fast by two
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next

        # Delete the middle node
        slow.next = slow.next.next

        return head

# Helper function to print the list
def print_list(head):
    while head:
        print(head.val, end=" -> ")
        head = head.next
    print("None")

# Example usage
head = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
solution = Solution()
new_head = solution.deleteMiddle(head)

# Print the updated list
print_list(new_head)  # Expected output: 1 -> 2 -> 4 -> 5 -> None
