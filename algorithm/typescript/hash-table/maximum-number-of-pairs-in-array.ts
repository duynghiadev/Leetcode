// Source: https://leetcode.com/problems/maximum-number-of-pairs-in-an-array/

function numberOfPairs(nums: number[]): number[] {
  const map = new Map<number, number>();

  // Count occurrences of each number
  for (const num of nums) {
    map.set(num, (map.get(num) || 0) + 1);
  }

  let totalPairs = 0;
  let totalLeft = 0;

  // Calculate total pairs and leftovers
  for (const value of map.values()) {
    totalPairs += Math.floor(value / 2);
  }

  totalLeft = nums.length - totalPairs * 2;

  return [totalPairs, totalLeft];
}

// Test cases
console.log(numberOfPairs([1, 3, 2, 1, 3, 2, 2])); // Output: [3, 1]
console.log(numberOfPairs([1, 1])); // Output: [1, 0]
console.log(numberOfPairs([0])); // Output: [0, 1]
console.log(numberOfPairs([1, 2, 3, 4])); // Output: [0, 4]
