# 1304. Find N Unique Integers Sum up to Zero

## Problem Description

Given an integer `n`, return any array containing `n` unique integers such that they add up to 0.

## Examples

### Example 1:
```
Input: n = 5
Output: [-7,-1,1,3,4]
Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].
```

### Example 2:
```
Input: n = 3
Output: [-1,0,1]
```

### Example 3:
```
Input: n = 1
Output: [0]
```

## Constraints

- `1 <= n <= 1000`

## Solution

The solution function `sumZero(n int) []int` is provided but left blank for you to implement.

## Approach

One simple approach is:
1. If `n` is odd, include 0 in the result and work with pairs for the remaining numbers
2. For pairs, add both positive and negative versions of the same number (e.g., 1 and -1)
3. This ensures the sum is always 0 and all numbers are unique
