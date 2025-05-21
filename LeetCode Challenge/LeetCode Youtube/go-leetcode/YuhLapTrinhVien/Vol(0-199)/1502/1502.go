package main

import "fmt"

/*
 * 📐 Arithmetic Progression Check Using Max-Min Formula
 * ================================================
 *
 * 🎯 Key Idea:
 * In an arithmetic progression, the difference between
 * any two consecutive numbers must be equal.
 *
 * 📝 Formula:
 * If a sequence is arithmetic, then:
 * (max - min) / (length - 1) = common difference
 *
 * 🌟 Examples:
 * 1. [2, 4, 6, 8]
 *    max = 8, min = 2, length = 4
 *    (8 - 2) / (4 - 1) = 2 (integer)
 *    Common difference is 2 ✅
 *
 * 2. [1, 2, 4]
 *    max = 4, min = 1, length = 3
 *    (4 - 1) / (3 - 1) = 1.5 (not integer)
 *    No common difference ❌
 */

func canMakeArithmeticProgression(arr []int) bool {
	// Handle small arrays
	if len(arr) <= 2 {
		return true
	}

	// Find min and max
	min, max := arr[0], arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// If all numbers are same, it's an arithmetic progression
	if min == max {
		return true
	}

	// Calculate expected difference
	n := len(arr)
	if (max-min)%(n-1) != 0 {
		return false // If not divisible evenly, can't be arithmetic
	}
	diff := (max - min) / (n - 1)

	// Use map to check if all required numbers exist
	seen := make(map[int]bool)
	for _, num := range arr {
		seen[num] = true
	}

	// Check if all numbers in sequence exist
	for i := 0; i < n; i++ {
		if !seen[min+i*diff] {
			return false
		}
	}

	return true
}

func main() {
	// Test cases
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))            // Output: true -> 4/2 = 2
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4}))            // Output: false -> 3/2 = 1.5
	fmt.Println(canMakeArithmeticProgression([]int{2, 4, 6, 8, 10, 12})) // Output: true -> 10/5 = 2
}
