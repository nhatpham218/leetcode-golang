# 1262. Greatest Sum Divisible by Three

## Problem Description

Given an integer array `nums`, return the maximum possible sum of elements of the array such that it is divisible by three.

## Examples

### Example 1:
```
Input: nums = [3,6,5,1,8]
Output: 18
Explanation: Pick numbers 3, 6, 1 and 8 their sum is 18 (maximum sum divisible by 3).
```

### Example 2:
```
Input: nums = [4]
Output: 0
Explanation: Since 4 is not divisible by 3, do not pick any number.
```

### Example 3:
```
Input: nums = [1,2,3,4,4]
Output: 12
Explanation: Pick numbers 1, 3, 4 and 4 their sum is 12 (maximum sum divisible by 3).
```

## Constraints

- `1 <= nums.length <= 4 * 10^4`
- `1 <= nums[i] <= 10^4`

## Solution

### Algorithm Overview

This problem requires finding the maximum sum of array elements that is divisible by 3.

### Approach

The key insight is to use dynamic programming to track the maximum sums with remainders 0, 1, and 2 when divided by 3.

### Complexity Analysis

- **Time Complexity**: O(n) where n is the length of nums
- **Space Complexity**: O(1)

### Implementation

```go
func maxSumDivThree(nums []int) int {
    // Find maximum sum divisible by 3
}
```

