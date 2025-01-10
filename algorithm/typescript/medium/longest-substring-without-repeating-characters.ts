// https://leetcode.com/problems/longest-substring-without-repeating-characters/

function lengthOfLongestSubstring(s: string): number {
  const seen = new Set<string>();
  let start = 0;
  let longest = 0;

  for (let end = 0; end < s.length; end++) {
    while (seen.has(s[end])) {
      seen.delete(s[start]);
      start++;
    }
    seen.add(s[end]);
    longest = Math.max(longest, end - start + 1);
  }

  return longest;
}

console.log(lengthOfLongestSubstring("abcabcbb")); // Output: 3
console.log(lengthOfLongestSubstring("bbbbb")); // Output: 1
console.log(lengthOfLongestSubstring("pwwkew")); // Output: 3
console.log(lengthOfLongestSubstring("")); // Output: 0
console.log(lengthOfLongestSubstring("abcdef")); // Output: 6
