# 1780. Check if Number is a Sum of Powers of Three

## Problem Description

Given an integer `n`, return `true` if it is possible to represent `n` as the sum of **distinct** powers of three. Otherwise, return `false`.

An integer `y` is a power of three if there exists an integer `x` such that `y == 3^x`.

## Examples

### Example 1:
```
Input: n = 12
Output: true
Explanation: 12 = 3^1 + 3^2 = 3 + 9
```

### Example 2:
```
Input: n = 91
Output: true
Explanation: 91 = 3^0 + 3^2 + 3^4 = 1 + 9 + 81
```

### Example 3:
```
Input: n = 21
Output: false
```

## Constraints

- `1 <= n <= 10^7`

## Solution Approach

The key insight is that if a number can be represented as a sum of distinct powers of three, then its base-3 representation should only contain digits 0 and 1. If any digit is 2, it means we would need to use the same power twice, which violates the "distinct powers" requirement.

### Algorithm:
1. Convert the number to base-3 representation
2. Check if any digit is 2
3. If no digit is 2, return true; otherwise, return false

### Time Complexity: O(n × log₃n)
### Space Complexity: O(n)

### Dynamic Programming Details

The DP approach works as follows:

1. **State Definition**: `dp[i]` = true if number `i` can be represented as a sum of distinct powers of 3
2. **State Transition**: For each power of 3, we can either use it or not use it
3. **Backward Iteration**: We iterate backwards to avoid using the same power multiple times
4. **Final State**: `dp[n]` gives us the answer

**DP Array Construction Example for n = 12:**
```
Powers of 3: [1, 3, 9]
Initial: dp[0] = true, others = false

After power 1: dp[1] = true
After power 3: dp[3] = true, dp[4] = true (4 = 1 + 3)
After power 9: dp[9] = true, dp[10] = true (10 = 1 + 9), dp[12] = true (12 = 3 + 9)
```

## Test Cases

The test file includes various test cases:
- Basic cases: 12, 91, 21
- Single powers: 1, 3, 9, 27
- Sums of powers: 4, 10
- Invalid cases: 2, 5
- Edge case: 0

## Solution Implementation

The solution is implemented in `1780.go` with the function signature:
```go
func checkPowersOfThree(n int) bool
```

### How the Solution Works

1. **Power Generation**: Generate all powers of 3 up to the input number n
2. **Dynamic Programming Array**: Create a boolean array `dp` where `dp[i]` represents whether number `i` can be represented as a sum of distinct powers of 3
3. **Base Case**: `dp[0] = true` (0 can be represented as an empty sum)
4. **State Transition**: For each power of 3, iterate backwards through the dp array to build larger numbers
5. **Result**: Return `dp[n]` which indicates whether n can be represented as a sum of distinct powers of 3

### Example Walkthrough

**Example 1: n = 12**
- Powers of 3 up to 12: [1, 3, 9]
- DP process:
  - Start with dp[0] = true
  - Use power 1: dp[1] = true
  - Use power 3: dp[3] = true, dp[4] = true (4 = 1 + 3)
  - Use power 9: dp[9] = true, dp[10] = true (10 = 1 + 9), dp[12] = true (12 = 3 + 9)
- Result: dp[12] = true ✓

**Example 2: n = 91**
- Powers of 3 up to 91: [1, 3, 9, 27, 81]
- DP process builds up to show 91 = 1 + 9 + 81
- Result: dp[91] = true ✓

**Example 3: n = 21**
- Powers of 3 up to 21: [1, 3, 9, 27]
- DP process shows no combination of distinct powers equals 21
- Result: dp[21] = false ✗

## Running the Tests

```bash
cd leetcode/1780.Check-if-Number-is-a-Sum-of-Powers-of-Three
go test -v
```

## Running the Solution

To test with custom input values, you can modify the test file or create a simple main function:

```go
package main

import "fmt"

func main() {
	testCases := []int{12, 91, 21, 1, 3, 9, 27, 4, 10, 2, 5, 0}
	
	for _, n := range testCases {
		result := checkPowersOfThree(n)
		fmt.Printf("checkPowersOfThree(%d) = %t\n", n, result)
	}
}
```
