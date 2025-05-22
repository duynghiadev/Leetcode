package main

import "fmt"

/*
 * 💰 Average Salary Calculator
 * ====================
 *
 * 📝 How it works:
 * 1. Find minimum and maximum salaries in the array
 * 2. Calculate sum of all salaries
 * 3. Subtract min and max from sum
 * 4. Divide by (n-2) to get average
 *
 * 🎯 Example: [4000, 3000, 1000, 2000]
 * Min: 1000 ⬇
 * Max: 4000 ⬆
 * Sum: 4000 + 3000 + 1000 + 2000 = 10000
 * Sum (excl. min & max): 10000 - 1000 - 4000 = 5000
 * Average: 5000 / 2 = 2500.0
 */
func average(salary []int) float64 {
	// Initialize min, max, and sum
	minSalary := salary[0]
	maxSalary := salary[0]
	sum := 0

	// Find min, max, and sum in one pass
	for _, s := range salary {
		if s < minSalary {
			minSalary = s // Update min
		}
		if s > maxSalary {
			maxSalary = s // Update max
		}
		sum += s // Add to sum
	}

	// Calculate average excluding min and max
	return float64(sum-minSalary-maxSalary) / float64(len(salary)-2)
}

func main() {
	// Test cases
	fmt.Println(average([]int{4000, 3000, 1000, 2000})) // Output: 2500.00000
	fmt.Println(average([]int{1000, 2000, 3000}))       // Output: 2000.00000
}
