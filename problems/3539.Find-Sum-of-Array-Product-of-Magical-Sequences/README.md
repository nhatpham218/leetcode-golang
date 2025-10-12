# 3539. Find Sum of Array Product of Magical Sequences

## Problem Description

You are given two integers, `m` and `k`, and an integer array `nums`.

A sequence of integers `seq` is called **magical** if:
- `seq` has a size of `m`.
- `0 <= seq[i] < nums.length`
- The binary representation of `2^seq[0] + 2^seq[1] + ... + 2^seq[m - 1]` has `k` set bits.

The **array product** of this sequence is defined as `prod(seq) = (nums[seq[0]] * nums[seq[1]] * ... * nums[seq[m - 1]])`.

Return the sum of the array products for all valid magical sequences.

Since the answer may be large, return it modulo `10^9 + 7`.

**Note**: A set bit refers to a bit in the binary representation of a number that has a value of 1.

## Examples

### Example 1:
```
Input: m = 5, k = 5, nums = [1,10,100,10000,1000000]
Output: 991600007
Explanation:
All permutations of [0, 1, 2, 3, 4] are magical sequences, each with an array product of 10^13.
```

### Example 2:
```
Input: m = 2, k = 2, nums = [5,4,3,2,1]
Output: 170
Explanation:
The magical sequences are [0, 1], [0, 2], [0, 3], [0, 4], [1, 0], [1, 2], [1, 3], [1, 4], [2, 0], [2, 1], [2, 3], [2, 4], [3, 0], [3, 1], [3, 2], [3, 4], [4, 0], [4, 1], [4, 2], and [4, 3].
```

### Example 3:
```
Input: m = 1, k = 1, nums = [28]
Output: 28
Explanation:
The only magical sequence is [0].
```

## Constraints

- `1 <= k <= m <= 30`
- `1 <= nums.length <= 50`
- `1 <= nums[i] <= 10^8`

## Solution

### Algorithm Overview

This problem requires finding all sequences of length `m` where the indices form a bitmask with exactly `k` set bits, then computing the sum of products for all such sequences.

### Key Insights

1. **Bitmask constraint**: The sum `2^seq[0] + 2^seq[1] + ... + 2^seq[m-1]` represents a bitmask where each element in the sequence corresponds to a bit position.
2. **Set bits**: We need exactly `k` set bits, meaning we need to choose `k` distinct positions from the available indices.
3. **Product calculation**: For each valid magical sequence, compute the product of corresponding elements from `nums`.
4. **Modular arithmetic**: Use modulo `10^9 + 7` for the final result.

### Approach

1. Generate all possible sequences of length `m` from indices `0` to `nums.length - 1`
2. For each sequence, check if the bitmask condition is satisfied (exactly `k` set bits)
3. Calculate the product of elements from `nums` at the positions specified by the sequence
4. Sum all products and return modulo `10^9 + 7`

### Complexity Analysis

- **Time Complexity**: O(n^m * m) where n is the length of nums and m is the sequence length
- **Space Complexity**: O(m) for recursion/iteration stack

### Implementation

```go
func magicalSum(m int, k int, nums []int) int {
    // Generate all sequences of length m
    // Check if sequence has exactly k set bits in bitmask
    // Calculate product for valid sequences
    // Return sum modulo 10^9 + 7
}
```

### Test Cases

- `m = 5, k = 5, nums = [1,10,100,10000,1000000]` → `991600007`
- `m = 2, k = 2, nums = [5,4,3,2,1]` → `170`
- `m = 1, k = 1, nums = [28]` → `28`

