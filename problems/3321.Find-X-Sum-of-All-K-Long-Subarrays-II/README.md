# 3321. Find X-Sum of All K-Long Subarrays II

## Problem Description

You are given an array `nums` of `n` integers and two integers `k` and `x`.

The **x-sum** of an array is calculated by the following procedure:

1. Count the occurrences of all elements in the array.
2. Keep only the occurrences of the top `x` most frequent elements. If two elements have the same number of occurrences, the element with the bigger value is considered more frequent.
3. Calculate the sum of the resulting array.

**Note** that if an array has less than `x` distinct elements, its x-sum is the sum of the array.

Return an integer array `answer` of length `n - k + 1` where `answer[i]` is the x-sum of the subarray `nums[i..i + k - 1]`.

## Examples

### Example 1:
```
Input: nums = [1,1,2,2,3,4,2,3], k = 6, x = 2
Output: [6,10,12]
Explanation:
- For subarray [1, 1, 2, 2, 3, 4], only elements 1 and 2 will be kept in the resulting array. Hence, answer[0] = 1 + 1 + 2 + 2.
- For subarray [1, 2, 2, 3, 4, 2], only elements 2 and 4 will be kept in the resulting array. Hence, answer[1] = 2 + 2 + 2 + 4. Note that 4 is kept in the array since it is bigger than 3 and 1 which occur the same number of times.
- For subarray [2, 2, 3, 4, 2, 3], only elements 2 and 3 are kept in the resulting array. Hence, answer[2] = 2 + 2 + 2 + 3 + 3.
```

### Example 2:
```
Input: nums = [3,8,7,8,7,5], k = 2, x = 2
Output: [11,15,15,15,12]
Explanation:
Since k == x, answer[i] is equal to the sum of the subarray nums[i..i + k - 1].
```

## Constraints

- `nums.length == n`
- `1 <= n <= 10^5`
- `1 <= nums[i] <= 10^9`
- `1 <= x <= k <= nums.length`

## Solution

### Algorithm Overview

This problem requires calculating the x-sum for each k-length subarray, where the x-sum is determined by selecting the top x most frequent elements (with ties broken by value).

### Key Insights

1. **Frequency Counting**: For each subarray, count the frequency of each element.
2. **Top X Selection**: Select the top x most frequent elements. If frequencies are equal, prefer larger values.
3. **Sum Calculation**: The x-sum includes all occurrences of the selected elements.
4. **Edge Case**: If there are fewer than x distinct elements, include all elements in the sum.

### Approach

1. Iterate through all possible subarrays of length k
2. For each subarray:
   - Count frequencies of all elements
   - Sort elements by frequency (descending), then by value (descending) for ties
   - Select top x elements
   - Calculate sum including all occurrences

### Complexity Analysis

- **Time Complexity**: O(n * k * log k) where n is array length
- **Space Complexity**: O(k) for frequency map

### Implementation

```go
func findXSum(nums []int, k int, x int) []int64 {
    // Implement sliding window with frequency tracking
}
```

### Test Cases

- `nums = [1,1,2,2,3,4,2,3], k = 6, x = 2` → `[6,10,12]`
- `nums = [3,8,7,8,7,5], k = 2, x = 2` → `[11,15,15,15,12]`



