# 🚀 LeetCode Practice Guide

## 🏆 Goal

Improve problem-solving skills and algorithmic thinking using **JavaScript** and **Go** while practicing LeetCode problems.

## 📌 Plan

1. **Choose a Problem** : Start with easy problems, then move to medium and hard ones.
2. **Understand the Problem** : Read the description carefully and identify constraints.
3. **Plan a Solution** : Think about brute force and optimize using better algorithms.
4. **Implement in JavaScript & Go** : Solve in both languages to strengthen your skills.
5. **Analyze Time & Space Complexity** : Optimize where possible.
6. **Review & Learn** : Read discussions and improve your solution.

## ⚡ Programming Languages

- **JavaScript** (for frontend and backend problem-solving)
- **Go** (for performance-oriented solutions)

## 🚀 Example Solutions

### 🟢 Two Sum (Easy)

#### JavaScript

```javascript
function twoSum(nums, target) {
  let map = new Map();
  for (let i = 0; i < nums.length; i++) {
    let complement = target - nums[i];
    if (map.has(complement)) {
      return [map.get(complement), i];
    }
    map.set(nums[i], i);
  }
  return [];
}
```

#### Go

```go
package main
import "fmt"

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        if j, ok := m[target-num]; ok {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    fmt.Println(twoSum(nums, target)) // Output: [0,1]
}
```

## 📅 Weekly Challenge

- Solve **2-3 problems per week**
- Focus on **one topic at a time** (Arrays, Strings, Recursion, DP, etc.)
- Try **both JavaScript & Go** for each problem

## 📚 Useful Resources

- [LeetCode Discussions](https://leetcode.com/discuss/)
- [My LeetCode](https://leetcode.com/u/duynghia22302/)
- [NeetCode Guide](https://neetcode.io/)
- [Big O Cheat Sheet](https://www.bigocheatsheet.com/)

## 🚀 Let's Code and Improve! 🔥
