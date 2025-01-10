package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

// addTwoNumbers adds two numbers represented by linked lists and returns the sum as a linked list.
func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	dummyHead := &Node{} // Dummy node to simplify list creation.
	current := dummyHead
	carry := 0

	// Traverse both lists until both are nil and there's no carry.
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Calculate the digit and carry.
		carry = sum / 10
		current.Next = &Node{Val: sum % 10}
		current = current.Next
	}

	return dummyHead.Next
}

// Helper function to create a linked list from a slice.
func createList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	head := &Node{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &Node{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to print a linked list.
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})

	// Expected output: 7 -> 0 -> 8 (342 + 465 = 807)
	result := addTwoNumbers(l1, l2)
	printList(result)
}
