/*
 * 🔢 Arithmetic Progression Checker
 * ==============================
 *
 * 📝 What is Arithmetic Progression?
 * - Each number differs from the next by a constant amount
 * - Example: [2, 4, 6, 8] → difference is always 2
 *
 * 🎯 Solution Steps:
 * 1. Find min and max numbers
 * 2. Calculate expected difference using (max-min)/(length-1)
 * 3. Check if all numbers in sequence exist
 */

function canMakeArithmeticProgression(arr: number[]): boolean {
  // Handle arrays with 0 or 1 elements
  if (arr.length <= 2) return true;

  // Find min and max
  let min = Math.min(...arr);
  let max = Math.max(...arr);

  // If all numbers are same
  if (min === max) return true;

  const n = arr.length;

  // Check if difference will be whole number
  if ((max - min) % (n - 1) !== 0) return false;

  // Calculate the expected difference
  const diff = (max - min) / (n - 1);

  // Use Set for O(1) lookup
  const numSet = new Set(arr);

  // Check if all numbers in sequence exist
  for (let i = 0; i < n; i++) {
    if (!numSet.has(min + i * diff)) {
      return false;
    }
  }

  return true;
}

// Test cases
console.log(canMakeArithmeticProgression([3, 5, 1])); // true: [1, 3, 5]
console.log(canMakeArithmeticProgression([1, 2, 4])); // false
console.log(canMakeArithmeticProgression([1, 1, 1, 1])); // true: all same
