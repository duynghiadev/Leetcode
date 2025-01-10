package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	var q []*TreeNode
	q = append(q, root)
	res := -1
	for len(q) > 0 {
		length := len(q)
		level := []int{}

		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		if len(level) > 0 {
			res = level[0]
		}
	}
	return res
}

func main() {
	// Example tree:
	//        2
	//       / \
	//      1   3
	// The bottom-left value is 1.

	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	result := findBottomLeftValue(root)
	fmt.Println("Bottom-left value:", result) // Output: 1
}
