# 1317. Convert Integer to the Sum of Two No-Zero Integers

**Difficulty:** Easy

## Problem Description

No-Zero integer is a positive integer that does not contain any 0 in its decimal representation.

Given an integer `n`, return a list of two integers `[a, b]` where:
- `a` and `b` are No-Zero integers.
- `a + b = n`

The test cases are generated so that there is at least one valid solution. If there are many valid solutions, you can return any of them.

## Examples

**Example 1:**
```
Input: n = 2
Output: [1,1]
Explanation: Let a = 1 and b = 1.
Both a and b are no-zero integers, and a + b = 2 = n.
```

**Example 2:**
```
Input: n = 11
Output: [2,9]
Explanation: Let a = 2 and b = 9.
Both a and b are no-zero integers, and a + b = 11 = n.
Note that there are other valid answers as [8, 3] that can be accepted.
```

## Constraints

- `2 <= n <= 10^4`

## Solution Approach

The solution uses a simple brute force approach:

1. Iterate through all possible values of `a` from 1 to `n-1`
2. For each `a`, calculate `b = n - a`
3. Check if both `a` and `b` are no-zero integers (don't contain digit 0)
4. Return the first valid pair found

### Helper Function

`hasNoZero(num int) bool`: Checks if a number contains no zero digits by examining each digit using modulo and division operations.

## Time Complexity

- **Time:** O(n * log(n)) where n is the input number. We iterate through at most n values, and for each value, we check if it has no zeros which takes O(log(n)) time.
- **Space:** O(1) - only using constant extra space.

## Implementation

```go
func getNoZeroIntegers(n int) []int {
    for a := 1; a < n; a++ {
        b := n - a
        if hasNoZero(a) && hasNoZero(b) {
            return []int{a, b}
        }
    }
    return nil
}

func hasNoZero(num int) bool {
    if num == 0 {
        return false
    }
    for num > 0 {
        if num%10 == 0 {
            return false
        }
        num /= 10
    }
    return true
}
```
