# 1458. Maximum Dot Product of Two Subsequences

## Problem Description

Given two arrays `nums1` and `nums2`.

Return the maximum dot product between **non-empty** subsequences of nums1 and nums2 with the same length.

A subsequence of a array is a new array which is formed from the original array by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, `[2,3,5]` is a subsequence of `[1,2,3,4,5]` while `[1,5,3]` is not).

## Examples

### Example 1:
```
Input: nums1 = [2,1,-2,5], nums2 = [3,0,-6]
Output: 18
Explanation: Take subsequence [2,-2] from nums1 and subsequence [3,-6] from nums2.
Their dot product is (2*3 + (-2)*(-6)) = 18.
```

### Example 2:
```
Input: nums1 = [3,-2], nums2 = [2,-6,7]
Output: 21
Explanation: Take subsequence [3] from nums1 and subsequence [7] from nums2.
Their dot product is (3*7) = 21.
```

### Example 3:
```
Input: nums1 = [-1,-1], nums2 = [1,1]
Output: -1
Explanation: Take subsequence [-1] from nums1 and subsequence [1] from nums2.
Their dot product is -1.
```

## Constraints

- `1 <= nums1.length, nums2.length <= 500`
- `-1000 <= nums1[i], nums2[i] <= 1000`

## Solution

### Algorithm Overview

This solution uses **Dynamic Programming** with space optimization using a rolling array technique.

### Approach

1. **Space Optimization**: Swap arrays if needed so that `nums1` is longer, allowing us to iterate over the shorter array in the inner loop.

2. **Rolling Array**: Use `dp[i&1][j]` where `i&1` alternates between 0 and 1, reducing space from O(n×m) to O(m).

3. **DP Transitions**: For each position (i, j), consider:
   - Skip current element from nums1
   - Skip current element from nums2
   - Take product nums1[i] * nums2[j] alone
   - Take product + best result from remaining elements

### Complexity Analysis

- **Time Complexity**: O(n × m) where n = len(nums1), m = len(nums2)
- **Space Complexity**: O(m) due to the rolling array optimization

