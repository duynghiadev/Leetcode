// time complexity: O(n^2) (nested loop)
function numIdenticalPairs(nums: number[]): number {
  let count: number = 0;
  const n: number = nums.length;

  for (let i = 0; i < n - 1; i++) {
    for (let j = i + 1; j < n; j++) {
      if (nums[i] === nums[j]) {
        count++;
      }
    }
  }
  return count;
}

// time complexity: O(n) (single pass, using map)
function numIdenticalPairs_1(nums: number[]): number {
  const freq: Map<number, number> = new Map();
  let count: number = 0;

  for (const num of nums) {
    count += freq.get(num) || 0;
    freq.set(num, (freq.get(num) || 0) + 1);
  }
  return count;
}

// code in video: O(n)
function numIdenticalPairs_2(nums: number[]): number {
  let result: number = 0;
  const numberCountMap: Map<number, number> = new Map();

  for (let i = 0; i < nums.length; i++) {
    const count = numberCountMap.get(nums[i]) || 0;
    if (count > 0) {
      result += count;
      numberCountMap.set(nums[i], count + 1);
    } else {
      numberCountMap.set(nums[i], 1);
    }
  }
  return result;
}

const nums: number[] = [1, 2, 3, 1, 1, 3];
console.log("Total Good Pair O(n^2):", numIdenticalPairs(nums));
console.log("Total Good Pair O(n):", numIdenticalPairs_1(nums));
console.log("Total Good Pair O(n):", numIdenticalPairs_2(nums));
