# 3350. Adjacent Increasing Subarrays Detection II

## Problem Description

Given an array `nums` of n integers, your task is to find the maximum value of `k` for which there exist two adjacent subarrays of length `k` each, such that both subarrays are strictly increasing. Specifically, check if there are two subarrays of length `k` starting at indices `a` and `b` (`a < b`), where:

- Both subarrays `nums[a..a + k - 1]` and `nums[b..b + k - 1]` are strictly increasing.
- The subarrays must be adjacent, meaning `b = a + k`.

Return the maximum possible value of `k`.

A subarray is a contiguous non-empty sequence of elements within an array.

## Examples

### Example 1:
```
Input: nums = [2,5,7,8,9,2,3,4,3,1]
Output: 3
Explanation:
The subarray starting at index 2 is [7, 8, 9], which is strictly increasing.
The subarray starting at index 5 is [2, 3, 4], which is also strictly increasing.
These two subarrays are adjacent, and 3 is the maximum possible value of k for which two such adjacent strictly increasing subarrays exist.
```

### Example 2:
```
Input: nums = [1,2,3,4,4,4,4,5,6,7]
Output: 2
Explanation:
The subarray starting at index 0 is [1, 2], which is strictly increasing.
The subarray starting at index 2 is [3, 4], which is also strictly increasing.
These two subarrays are adjacent, and 2 is the maximum possible value of k for which two such adjacent strictly increasing subarrays exist.
```

## Constraints

- `2 <= nums.length <= 2 * 10^5`
- `-10^9 <= nums[i] <= 10^9`

## Solution

### Algorithm Overview

The problem requires finding the maximum length `k` for which two adjacent subarrays of length `k` are both strictly increasing.

### Key Insights

1. **Adjacent subarrays**: If the first subarray starts at index `i`, the second must start at index `i + k`.
2. **Strictly increasing**: Each element in a subarray must be greater than the previous element.
3. **Maximum k**: We need to find the largest possible value of `k` that satisfies the conditions.

### Approach

1. Iterate through all possible starting positions for pairs of adjacent subarrays
2. For each position, determine the maximum length `k` where both subarrays are strictly increasing
3. Track and return the maximum `k` found

### Complexity Analysis

- **Time Complexity**: O(n^2) where n is the array length
- **Space Complexity**: O(1) for the algorithm itself

### Implementation

```go
func maxIncreasingSubarrays(nums []int) int {
    // Find the maximum k for two adjacent strictly increasing subarrays
}
```

### Test Cases

- `nums = [2,5,7,8,9,2,3,4,3,1]` → `3`
- `nums = [1,2,3,4,4,4,4,5,6,7]` → `2`

