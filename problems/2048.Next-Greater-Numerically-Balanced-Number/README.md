# 2048. Next Greater Numerically Balanced Number

**Difficulty:** Medium

## Problem Description

An integer `x` is **numerically balanced** if for every digit `d` in the number `x`, there are **exactly** `d` occurrences of that digit in `x`.

Given an integer `n`, return the **smallest numerically balanced** number **strictly greater** than `n`.

## Examples

### Example 1:
```
Input: n = 1
Output: 22
Explanation: 
22 is numerically balanced since:
- The digit 2 occurs 2 times. 
It is also the smallest numerically balanced number strictly greater than 1.
```

### Example 2:
```
Input: n = 1000
Output: 1333
Explanation: 
1333 is numerically balanced since:
- The digit 1 occurs 1 time.
- The digit 3 occurs 3 times. 
It is also the smallest numerically balanced number strictly greater than 1000.
Note that 1022 cannot be the answer because 0 appeared more than 0 times.
```

### Example 3:
```
Input: n = 3000
Output: 3133
Explanation: 
3133 is numerically balanced since:
- The digit 1 occurs 1 time.
- The digit 3 occurs 3 times.
It is also the smallest numerically balanced number strictly greater than 3000.
```

## Constraints

- `0 <= n <= 10^6`

## Solution Approach

The solution uses a pre-computation and binary search approach:

1. **Pre-generation (init)**: Generate all numerically balanced numbers up to 7 digits:
   - Use DFS to select combinations of digits 1-7
   - For each selection, generate all permutations
   - Store unique numbers in a sorted array

2. **Binary Search**: Use `sort.Search` to find the smallest number greater than `n`

### Key Insights

- Digits 1-7 are selected because digit `d` must appear `d` times
- For 7-digit numbers, max is 7 copies of digit 7 (7777777), but this exceeds 10^6
- Valid combinations are limited (e.g., {1}, {2,2}, {1,3,3,3}, {1,4,4,4,4})
- Pre-generating all valid numbers makes lookup O(log n)

### Time Complexity

- **Init**: O(P) where P is total permutations generated (one-time cost)
- **Query**: O(log N) where N is count of balanced numbers (~80 numbers)

### Space Complexity

- **O(N)** to store pre-generated balanced numbers

## Examples of Numerically Balanced Numbers

- `1`: digit 1 appears 1 time ✓
- `22`: digit 2 appears 2 times ✓
- `122`: digit 1 appears 1 time, digit 2 appears 2 times ✓
- `1333`: digit 1 appears 1 time, digit 3 appears 3 times ✓
- `5555`: digit 5 appears 5 times ✓
- `22333`: digit 2 appears 2 times, digit 3 appears 3 times ✓

