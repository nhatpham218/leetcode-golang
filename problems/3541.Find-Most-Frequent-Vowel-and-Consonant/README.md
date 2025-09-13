# 3541. Find Most Frequent Vowel and Consonant

## Problem Description

You are given a string `s` consisting of lowercase English letters ('a' to 'z').

Your task is to:

1. Find the vowel (one of 'a', 'e', 'i', 'o', or 'u') with the maximum frequency.
2. Find the consonant (all other letters excluding vowels) with the maximum frequency.
3. Return the sum of the two frequencies.

**Note**: If multiple vowels or consonants have the same maximum frequency, you may choose any one of them. If there are no vowels or no consonants in the string, consider their frequency as 0.

The frequency of a letter x is the number of times it occurs in the string.

## Examples

### Example 1:
```
Input: s = "successes"
Output: 6
Explanation:
The vowels are: 'u' (frequency 1), 'e' (frequency 2). The maximum frequency is 2.
The consonants are: 's' (frequency 4), 'c' (frequency 2). The maximum frequency is 4.
The output is 2 + 4 = 6.
```

### Example 2:
```
Input: s = "aeiaeia"
Output: 3
Explanation:
The vowels are: 'a' (frequency 3), 'e' (frequency 2), 'i' (frequency 2). The maximum frequency is 3.
There are no consonants in s. Hence, maximum consonant frequency = 0.
The output is 3 + 0 = 3.
```

## Constraints

- `1 <= s.length <= 100`
- `s` consists of lowercase English letters only.

## Solution

### Algorithm Overview

This is a **frequency counting** problem that requires tracking character frequencies separately for vowels and consonants.

### Key Insights

1. **Separate tracking**: Count frequencies of vowels and consonants separately
2. **Maximum frequency**: Find the maximum frequency in each category
3. **Handle empty categories**: If no vowels or consonants exist, their max frequency is 0

### Approach

1. Create frequency maps for vowels and consonants
2. Iterate through the string and count frequencies
3. Find maximum frequency in each category
4. Return the sum of both maximum frequencies

### Complexity Analysis

- **Time Complexity**: O(n) - Single pass through the string
- **Space Complexity**: O(1) - At most 26 characters to track

### Implementation

```go
func maxFreqSum(s string) int {
    // Count frequencies of vowels and consonants separately
    // Find maximum frequency in each category
    // Return sum of both maximums
}
```

### Test Cases

- `"successes"` → `6` (vowel max: 2, consonant max: 4)
- `"aeiaeia"` → `3` (vowel max: 3, consonant max: 0)
- `"a"` → `1` (vowel max: 1, consonant max: 0)
- `"b"` → `1` (vowel max: 0, consonant max: 1)
- `"aeiou"` → `1` (all unique vowels)
- `"bcdfg"` → `1` (all unique consonants)
