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
public class AverageSalary {
	public static double average(int[] salary) {
		// Initialize min, max, and sum
		int minSalary = salary[0];
		int maxSalary = salary[0];
		long sum = 0; // Use long to avoid overflow

		// Find min, max, and sum in one pass
		for (int s : salary) {
			if (s < minSalary) {
				minSalary = s; // Update min
			}
			if (s > maxSalary) {
				maxSalary = s; // Update max
			}
			sum += s; // Add to sum
		}

		// Calculate average excluding min and max
		return (double) (sum - minSalary - maxSalary) / (salary.length - 2);
	}

	public static void main(String[] args) {
		// Test cases
		System.out.println(average(new int[] { 4000, 3000, 1000, 2000 })); // Output: 2500.00000
		System.out.println(average(new int[] { 1000, 2000, 3000 })); // Output: 2000.00000
	}
}