/*
 * 🔢 Arithmetic Progression Checker
 * ==============================
 *
 * 📝 How it works:
 * 1. [1, 3, 5] is valid because:
 *    - min = 1, max = 5, length = 3
 *    - (5-1)/(3-1) = 2 (whole number)
 *    - Numbers needed: 1, 1+2, 1+2+2 → All exist!
 *
 * 2. [1, 2, 4] is not valid because:
 *    - min = 1, max = 4, length = 3
 *    - (4-1)/(3-1) = 1.5 (not whole number)
 */

class MeticProgression {
	public boolean canMakeArithmeticProgression(int[] arr) {
		// Handle small arrays
		if (arr.length <= 2)
			return true;

		// Find min and max
		int min = arr[0];
		int max = arr[0];
		for (int num : arr) {
			min = Math.min(min, num);
			max = Math.max(max, num);
		}

		// If all numbers are same
		if (min == max)
			return true;

		int n = arr.length;

		// Check if difference will be whole number
		if ((max - min) % (n - 1) != 0)
			return false;

		// Calculate expected difference
		int diff = (max - min) / (n - 1);

		// Use boolean array for O(1) lookup
		boolean[] seen = new boolean[n];

		// Mark each number's position
		for (int num : arr) {
			// Calculate expected position
			int pos = (num - min) / diff;

			// Check if position is valid and not already seen
			if (pos < 0 || pos >= n || seen[pos])
				return false;
			if ((num - min) % diff != 0)
				return false;

			seen[pos] = true;
		}

		return true;
	}

	// Test the solution
	public static void main(String[] args) {
		MeticProgression solution = new MeticProgression();
		System.out.println(solution.canMakeArithmeticProgression(new int[] { 3, 5, 1 })); // true
		System.out.println(solution.canMakeArithmeticProgression(new int[] { 1, 2, 4 })); // false
		System.out.println(solution.canMakeArithmeticProgression(new int[] { 1, 1, 1 })); // true
	}
}