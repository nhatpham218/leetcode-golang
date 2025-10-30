# 1526. Minimum Number of Increments on Subarrays to Form a Target Array

## Problem Description

You are given an integer array target. You have an integer array initial of the same size as target with all elements initially zeros.

In one operation you can choose any subarray from initial and increment each value by one.

Return the minimum number of operations to form a target array from initial.

The test cases are generated so that the answer fits in a 32-bit integer.

## Examples

### Example 1:
```
Input: target = [1,2,3,2,1]
Output: 3
Explanation: We need at least 3 operations to form the target array from the initial array.
[0,0,0,0,0] increment 1 from index 0 to 4 (inclusive).
[1,1,1,1,1] increment 1 from index 1 to 3 (inclusive).
[1,2,2,2,1] increment 1 at index 2.
[1,2,3,2,1] target array is formed.
```

### Example 2:
```
Input: target = [3,1,1,2]
Output: 4
Explanation: [0,0,0,0] -> [1,1,1,1] -> [1,1,1,2] -> [2,1,1,2] -> [3,1,1,2]
```

### Example 3:
```
Input: target = [3,1,5,4,2]
Output: 7
Explanation: [0,0,0,0,0] -> [1,1,1,1,1] -> [2,1,1,1,1] -> [3,1,1,1,1] -> [3,1,2,2,2] -> [3,1,3,3,2] -> [3,1,4,4,2] -> [3,1,5,4,2].
```

## Constraints

- `1 <= target.length <= 10^5`
- `1 <= target[i] <= 10^5`

## Solution

### Algorithm Overview

This problem requires finding the minimum number of operations to transform an array of zeros into the target array, where each operation increments all elements in a contiguous subarray by 1.

### Approach

The key insight is that we need to count the number of "increasing steps" in the array. When we move from one element to the next:
- If the next element is larger, we need additional operations equal to the difference
- If the next element is smaller or equal, we don't need any new operations

### Key Points

- Start with the first element (it requires target[0] operations from 0)
- For each subsequent element, only add operations if it's larger than the previous
- The total operations equals the sum of all positive differences

### Complexity Analysis

- **Time Complexity**: O(n)
  - Single pass through the array

- **Space Complexity**: O(1)
  - Only constant extra space is used

### Implementation

```go
func minNumberOperations(target []int) int {

}
```

### Test Cases

- target = [1,2,3,2,1] -> Output: 3
- target = [3,1,1,2] -> Output: 4
- target = [3,1,5,4,2] -> Output: 7

