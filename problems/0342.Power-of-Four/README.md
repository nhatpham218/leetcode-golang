# 342. Power of Four

## Problem Description

Given an integer `n`, return `true` if it is a power of four. Otherwise, return `false`.

An integer `n` is a power of four, if there exists an integer `x` such that `n == 4^x`.

## Examples

### Example 1:
```
Input: n = 16
Output: true
Explanation: 16 = 4^2
```

### Example 2:
```
Input: n = 5
Output: false
Explanation: 5 is not a power of 4
```

### Example 3:
```
Input: n = 1
Output: true
Explanation: 1 = 4^0
```

## Solution Approach

### Method 1: Mathematical Approach
A number is a power of four if and only if:
1. It's positive (greater than 0)
2. It's a power of 2 (only one bit is set in binary representation)
3. The set bit is at an even position (0, 2, 4, 6, ...)

**Key Insight**: For powers of 4, the number minus 1 is divisible by 3.
- 4^1 - 1 = 3 (divisible by 3)
- 4^2 - 1 = 15 (divisible by 3)
- 4^3 - 1 = 63 (divisible by 3)
- 4^4 - 1 = 255 (divisible by 3)

**Algorithm**:
1. Check if n > 0
2. Check if n is a power of 2 using `n & (n-1) == 0`
3. Check if `(n-1) % 3 == 0`

### Method 2: Bit Manipulation
Use a bit mask to check if the set bit is at an even position:
- `0x55555555` in binary is `01010101...` (1s at even positions)
- If `n & 0x55555555 != 0`, then the set bit is at an even position

## Time and Space Complexity

- **Time Complexity**: O(1) - constant time operations
- **Space Complexity**: O(1) - no extra space used

## Test Cases

The solution includes comprehensive test cases covering:
- Powers of 4: 1, 4, 16, 64, 256, 1024
- Powers of 2 that are not powers of 4: 2, 8, 32, 128
- Non-powers of 2: 3, 5, 6, 7, 9, 10, 12, 15, 17
- Edge cases: 0, negative numbers
- Large numbers: 1073741824 (4^15)

## Running the Code

### Run the main program:
```bash
go run 0342.go
```

### Run the tests:
```bash
go test
```

### Run benchmarks:
```bash
go test -bench=.
```

## Notes

- Both methods give identical results
- The mathematical approach is more intuitive and easier to understand
- The bit manipulation approach might be slightly faster in practice
- The solution handles all edge cases including negative numbers and zero
