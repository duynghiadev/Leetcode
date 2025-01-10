package main

import "fmt"

func combine(n int, k int) [][]int {
	result := [][]int{}
	current := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		// If the current combination is of the required size, add it to the result
		if len(current) == k {
			temp := make([]int, k)
			copy(temp, current)
			result = append(result, temp)
			return
		}

		// Iterate through the numbers from `start` to `n`
		for i := start; i <= n; i++ {
			current = append(current, i)       // Choose
			backtrack(i + 1)                   // Explore
			current = current[:len(current)-1] // Un-choose
		}
	}

	backtrack(1) // Start backtracking from 1
	return result
}

func main() {
	n := 4
	k := 2
	result := combine(n, k)
	fmt.Println(result)
}
