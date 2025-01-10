package main

import "fmt"

// customStack is a simple stack implementation.
type customStack struct {
	stack []int
}

// Push adds an element to the top of the stack.
func (cs *customStack) Push(val int) {
	cs.stack = append(cs.stack, val)
}

// Pop removes the top element from the stack.
func (cs *customStack) Pop() {
	if len(cs.stack) > 0 {
		cs.stack = cs.stack[:len(cs.stack)-1]
	}
}

// Front returns the top element of the stack.
func (cs *customStack) Front() (int, bool) {
	if len(cs.stack) > 0 {
		return cs.stack[len(cs.stack)-1], true
	}
	return 0, false
}

// Size returns the size of the stack.
func (cs *customStack) Size() int {
	return len(cs.stack)
}

// max returns the maximum height in the range [start, end].
func max(height []int, start, end int) int {
	maxVal := 0
	for i := start; i <= end; i++ {
		if height[i] > maxVal {
			maxVal = height[i]
		}
	}
	return maxVal
}

// trap calculates the amount of water that can be trapped.
func trap(height []int) int {
	output := 0
	customStack := customStack{
		stack: make([]int, 0),
	}

	customStack.Push(0)

	for i := 1; i < len(height); i++ {
		for customStack.Size() != 0 {
			front, _ := customStack.Front()
			if height[front] <= height[i] {
				output += (i - front - 1) * (height[front] - max(height, front+1, i-1))
				customStack.Pop()
			} else {
				output += (i - front - 1) * (height[i] - max(height, front+1, i-1))
				break
			}
		}
		customStack.Push(i)
	}
	return output
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height)) // Output: 6
}
