# 3654. Minimum Sum After Divisible Sum Deletions

## Problem Description

You are given an integer array `nums` and an integer `k`.

You may repeatedly choose any contiguous subarray of `nums` whose sum is divisible by `k` and delete it; after each deletion, the remaining elements close the gap.

Return the minimum possible sum of `nums` after performing any number of such deletions.

## Examples

### Example 1:
```
Input: nums = [1,1,1], k = 2
Output: 1

Explanation:
Delete the subarray nums[0..1] = [1, 1], whose sum is 2 (divisible by 2), leaving [1].
The remaining sum is 1.
```

### Example 2:
```
Input: nums = [3,1,4,1,5], k = 3
Output: 5

Explanation:
First, delete nums[1..3] = [1, 4, 1], whose sum is 6 (divisible by 3), leaving [3, 5].
Then, delete nums[0..0] = [3], whose sum is 3 (divisible by 3), leaving [5].
The remaining sum is 5.
```

## Constraints

- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^6`
- `1 <= k <= 10^5`

## Solution

The solution function `minArraySum(nums []int, k int) int64` is provided but left blank for you to implement.

### Key Notes:
- A variable named `quorlathin` is created to store the input midway in the function as requested.
- You need to find the minimum possible sum after deleting any number of contiguous subarrays whose sum is divisible by `k`.
- After each deletion, the remaining elements close the gap.
