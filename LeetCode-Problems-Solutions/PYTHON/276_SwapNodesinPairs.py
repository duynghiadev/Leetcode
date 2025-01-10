# Author name:Chengalva Sai Harikha
# Date:15/10/2022
# Time complexity:O(n) where n is the length
# Space Complexity:O(1)
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        if not head or not head.next:
            return head

        previous, current = None, head
        answer = head.next  # New head will be the second node

        while current and current.next:
            adj = current.next
            if previous:
                previous.next = adj
            current.next = adj.next
            adj.next = current

            previous = current
            current = current.next

        return answer

# Helper function to print the linked list
def print_list(head: ListNode):
    current = head
    while current:
        print(current.val, end=" -> ")
        current = current.next
    print("None")

# Helper function to create a linked list from a list of values
def create_linked_list(values):
    head = ListNode(values[0])
    current = head
    for value in values[1:]:
        current.next = ListNode(value)
        current = current.next
    return head

# Test the function
values = [1, 2, 3, 4]  # Input list
head = create_linked_list(values)
print("Original list:")
print_list(head)

solution = Solution()
swapped_head = solution.swapPairs(head)

print("List after swapping pairs:")
print_list(swapped_head)
