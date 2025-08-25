# 2787. Ways to Express an Integer as Sum of Powers

## Problem Description

Given two positive integers `n` and `x`, return the number of different ways to express `n` as the sum of unique powers of `x`.

## Examples

### Example 1:
- Input: n = 10, x = 2
- Output: 1
- Explanation: 10 = 3² + 1² = 9 + 1

### Example 2:
- Input: n = 4, x = 1
- Output: 2
- Explanation: 4 = 4¹ or 4 = 1¹ + 1¹ + 1¹ + 1¹

### Example 3:
- Input: n = 1, x = 1
- Output: 1
- Explanation: 1 = 1¹

## Constraints

- 1 <= n <= 300
- 1 <= x <= 5

## Function Signature

```go
func numberOfWays(n int, x int) int
```

## Solution Approach

### Dynamic Programming Solution

This problem is solved using a **dynamic programming** approach similar to the classic "coin change" problem, but with powers of x instead of coins.

#### Algorithm Steps:

1. **Initialize DP array**: `dp[i]` represents the number of ways to express integer `i` as a sum of powers of `x`
   - `dp[0] = 1` (base case: empty sum)

2. **Generate powers**: For each power value from 1 to n:
   - Calculate `power^x` (e.g., if x=2, calculate 1², 2², 3², etc.)
   - Stop when the power value exceeds n

3. **Update DP array**: For each power value `val`:
   - Iterate backwards from n down to `val`
   - Update `dp[i] = (dp[i] + dp[i-val]) % mod`
   - This adds the contribution of using the current power value

4. **Return result**: `dp[n]` contains the final answer

#### Key Insights:

- **Unique combinations**: Each power can only be used once in a combination
- **Order doesn't matter**: 3² + 1² is the same as 1² + 3²
- **Modulo arithmetic**: Use modulo 10^9 + 7 to handle large numbers
- **Backward iteration**: Prevents counting the same power multiple times

#### Time Complexity: O(n²)
#### Space Complexity: O(n)

## Example Walkthrough

For n = 10, x = 2:

1. **Power 1**: 1² = 1
   - Update dp[1] = 1, dp[2] = 1, dp[3] = 1, ..., dp[10] = 1

2. **Power 2**: 2² = 4
   - Update dp[4] = 1, dp[5] = 1, dp[6] = 1, ..., dp[10] = 1

3. **Power 3**: 3² = 9
   - Update dp[9] = 1, dp[10] = 1

4. **Power 4**: 4² = 16 > 10, stop

Result: dp[10] = 1 (only one way: 3² + 1²)

## Notes

- Each power can only be used once in a combination
- The order of powers in the sum doesn't matter
- We need to find unique combinations, not permutations
- The solution handles the modulo requirement for large numbers
