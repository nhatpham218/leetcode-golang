# 2273. Find Resultant Array After Removing Anagrams

## Problem Description

You are given a 0-indexed string array `words`, where `words[i]` consists of lowercase English letters.

In one operation, select any index `i` such that `0 < i < words.length` and `words[i - 1]` and `words[i]` are anagrams, and delete `words[i]` from `words`. Keep performing this operation as long as you can select an index that satisfies the conditions.

Return `words` after performing all operations. It can be shown that selecting the indices for each operation in any arbitrary order will lead to the same result.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase using all the original letters exactly once. For example, "dacb" is an anagram of "abdc".

## Examples

### Example 1:
```
Input: words = ["abba","baba","bbaa","cd","cd"]
Output: ["abba","cd"]
Explanation:
One of the ways we can obtain the resultant array is by using the following operations:
- Since words[2] = "bbaa" and words[1] = "baba" are anagrams, we choose index 2 and delete words[2].
  Now words = ["abba","baba","cd","cd"].
- Since words[1] = "baba" and words[0] = "abba" are anagrams, we choose index 1 and delete words[1].
  Now words = ["abba","cd","cd"].
- Since words[2] = "cd" and words[1] = "cd" are anagrams, we choose index 2 and delete words[2].
  Now words = ["abba","cd"].
We can no longer perform any operations, so ["abba","cd"] is the final answer.
```

### Example 2:
```
Input: words = ["a","b","c","d","e"]
Output: ["a","b","c","d","e"]
Explanation:
No two adjacent strings in words are anagrams of each other, so no operations are performed.
```

## Constraints

- `1 <= words.length <= 100`
- `1 <= words[i].length <= 10`
- `words[i]` consists of lowercase English letters.

## Solution

### Algorithm Overview

The problem requires removing consecutive anagrams from the array. An efficient approach is to iterate through the array and only keep words that are not anagrams of their predecessor.

### Key Insights

1. **Anagram detection**: Two words are anagrams if they have the same character frequency (or sorted characters).
2. **Single pass**: We can solve this in a single pass by comparing each word with the last kept word.
3. **Order matters**: Only adjacent anagrams need to be removed.

### Approach

1. Iterate through the array
2. Keep the first word in the result
3. For each subsequent word, check if it's an anagram of the last kept word
4. If not an anagram, add it to the result
5. If it's an anagram, skip it

### Complexity Analysis

- **Time Complexity**: O(n * m log m) where n is the number of words and m is the average word length
- **Space Complexity**: O(n * m) for the result array

### Implementation

```go
func removeAnagrams(words []string) []string {
    // Iterate through words
    // Check if current word is anagram of previous kept word
    // Keep non-anagram words
}
```

### Test Cases

- `words = ["abba","baba","bbaa","cd","cd"]` → `["abba","cd"]`
- `words = ["a","b","c","d","e"]` → `["a","b","c","d","e"]`

