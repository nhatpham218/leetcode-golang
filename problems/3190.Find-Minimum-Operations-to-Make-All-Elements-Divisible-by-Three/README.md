# 3190. Find Minimum Operations to Make All Elements Divisible by Three

## Problem Description

You are given an integer array `nums`. In one operation, you can add or subtract 1 from any element of `nums`.

Return the minimum number of operations to make all elements of `nums` divisible by 3.

## Examples

### Example 1:
```
Input: nums = [1,2,3,4]
Output: 3
Explanation:
All array elements can be made divisible by 3 using 3 operations:
- Subtract 1 from 1.
- Add 1 to 2.
- Subtract 1 from 4.
```

### Example 2:
```
Input: nums = [3,6,9]
Output: 0
```

## Constraints

- `1 <= nums.length <= 50`
- `1 <= nums[i] <= 50`

## Solution

### Algorithm Overview

This problem can be solved by checking each element's remainder when divided by 3.

### Approach

For each element in the array:
- If the element is already divisible by 3 (remainder = 0), no operation is needed
- If the remainder is 1, subtract 1 (1 operation)
- If the remainder is 2, add 1 (1 operation)

### Complexity Analysis

- **Time Complexity**: O(n) where n is the length of nums
- **Space Complexity**: O(1)

### Implementation

```go
func minimumOperations(nums []int) int {
    // Count operations needed for each element
}
```

