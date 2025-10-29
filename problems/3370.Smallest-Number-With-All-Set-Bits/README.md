# 3370. Smallest Number With All Set Bits

## Problem Description

You are given a positive number n.

Return the smallest number x greater than or equal to n, such that the binary representation of x contains only set bits.

## Examples

### Example 1:
```
Input: n = 5
Output: 7
Explanation:
The binary representation of 7 is "111".
```

### Example 2:
```
Input: n = 10
Output: 15
Explanation:
The binary representation of 15 is "1111".
```

### Example 3:
```
Input: n = 3
Output: 3
Explanation:
The binary representation of 3 is "11".
```

## Constraints

- `1 <= n <= 1000`

## Solution

### Algorithm Overview

This problem requires finding the smallest number greater than or equal to n where all bits in the binary representation are set to 1.

### Approach

1. **Understand Set Bits**: A number with all set bits has the form 2^k - 1 (e.g., 1, 3, 7, 15, 31, ...).
2. **Find the Smallest**: Starting from n, find the smallest number of the form 2^k - 1 that is >= n.
3. **Bit Manipulation**: We can find the position of the most significant bit in n and check if 2^k - 1 is enough.

### Key Points

- Numbers with all set bits: 1 (2^1-1), 3 (2^2-1), 7 (2^3-1), 15 (2^4-1), 31 (2^5-1), etc.
- If n itself is a number with all set bits, return n
- Otherwise, find the next number with all set bits

### Complexity Analysis

- **Time Complexity**: O(log n)
  - We only need to check the number of bits in n

- **Space Complexity**: O(1)
  - Only constant extra space is used

### Implementation

```go
func smallestNumber(n int) int {
	
}
```

### Test Cases

- n = 5 -> Output: 7
- n = 10 -> Output: 15
- n = 3 -> Output: 3
- n = 1 -> Output: 1
- Edge cases with powers of 2 minus 1

