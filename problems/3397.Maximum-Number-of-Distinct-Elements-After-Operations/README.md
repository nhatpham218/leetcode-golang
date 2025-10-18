# 3397. Maximum Number of Distinct Elements After Operations

## Problem Description

You are given an integer array `nums` and an integer `k`.

You are allowed to perform the following operation on each element of the array **at most once**:

- Add an integer in the range `[-k, k]` to the element.

Return the **maximum** possible number of **distinct** elements in `nums` after performing the operations.

## Examples

### Example 1:
```
Input: nums = [1,2,2,3,3,4], k = 2
Output: 6
Explanation:
nums changes to [-1, 0, 1, 2, 3, 4] after performing operations on the first four elements.
```

### Example 2:
```
Input: nums = [4,4,4,4], k = 1
Output: 3
Explanation:
By adding -1 to nums[0] and 1 to nums[1], nums changes to [3, 5, 4, 4].
```

## Constraints

- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^9`
- `0 <= k <= 10^9`

## Solution

### Algorithm Overview

This problem requires maximizing the number of distinct elements after applying operations. Each element can be modified by adding a value in the range `[-k, k]`.

### Key Insights

1. **Greedy approach**: To maximize distinct elements, we should try to assign each element a unique value within its possible range.
2. **Sorting**: Sorting the array helps in making greedy choices for each element.
3. **Range consideration**: For each element `nums[i]`, the possible values are `[nums[i] - k, nums[i] + k]`.

### Approach

1. Sort the array to process elements in order
2. For each element, try to assign it the smallest available value within its range `[nums[i] - k, nums[i] + k]`
3. Keep track of the last assigned value to avoid duplicates
4. Count the total number of distinct values assigned

### Complexity Analysis

- **Time Complexity**: O(n log n) where n is the length of nums (due to sorting)
- **Space Complexity**: O(1) or O(n) depending on the sorting implementation

### Implementation

```go
func maxDistinctElements(nums []int, k int) int {
    // Sort the array
    // For each element, try to assign the smallest available value
    // Keep track of last assigned value
    // Return count of distinct elements
}
```

### Test Cases

- `nums = [1,2,2,3,3,4], k = 2` → `6`
- `nums = [4,4,4,4], k = 1` → `3`

