package main

import (
	"fmt"
)

// time complexity: O(n^2) (nested loop)
func numIdenticalPairs(nums []int) int {
	count := 0
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

// time complexity: O(n) (single pass, using map)
func numIdenticalPairs_1(nums []int) int {
	freq := make(map[int]int)
	count := 0

	for _, num := range nums {
		count += freq[num]
		freq[num]++
	}
	return count
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	fmt.Println("Total Good Pairs O(n^2):", numIdenticalPairs(nums))
	fmt.Println("Total Good Pairs O(n):", numIdenticalPairs_1(nums))
}
