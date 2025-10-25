# 1716. Calculate Money in Leetcode Bank

## Problem Description

Hercy wants to save money for his first car. He puts money in the Leetcode bank every day.

He starts by putting in $1 on Monday, the first day. Every day from Tuesday to Sunday, he will put in $1 more than the day before. On every subsequent Monday, he will put in $1 more than the previous Monday.

Given `n`, return the total amount of money he will have in the Leetcode bank at the end of the nth day.

## Examples

### Example 1:
```
Input: n = 4
Output: 10
Explanation: After the 4th day, the total is 1 + 2 + 3 + 4 = 10.
```

### Example 2:
```
Input: n = 10
Output: 37
Explanation: After the 10th day, the total is (1 + 2 + 3 + 4 + 5 + 6 + 7) + (2 + 3 + 4) = 37. Notice that on the 2nd Monday, Hercy only puts in $2.
```

### Example 3:
```
Input: n = 20
Output: 96
Explanation: After the 20th day, the total is (1 + 2 + 3 + 4 + 5 + 6 + 7) + (2 + 3 + 4 + 5 + 6 + 7 + 8) + (3 + 4 + 5 + 6 + 7 + 8) = 96.
```

## Constraints

- `1 <= n <= 1000`

## Solution

### Algorithm Overview

The problem follows a pattern where:
- Week 1: starts with $1 on Monday, increases by $1 each day (1, 2, 3, 4, 5, 6, 7)
- Week 2: starts with $2 on Monday, increases by $1 each day (2, 3, 4, 5, 6, 7, 8)
- Week k: starts with $k on Monday, increases by $1 each day (k, k+1, k+2, ..., k+6)

### Approach

1. **Calculate complete weeks**: Determine how many complete 7-day weeks fit in n days.
2. **Sum complete weeks**: For each complete week k (1-indexed), the sum is: k + (k+1) + (k+2) + ... + (k+6) = 7k + 21
3. **Handle remaining days**: For the partial week, calculate the sum of the remaining days.

**Alternative Approach - Simulation:**
- Iterate through each day from 1 to n
- Track the current week number and day of the week
- Calculate the amount for each day based on the pattern

### Complexity Analysis

**Mathematical Approach:**
- **Time Complexity**: O(1)
  - Constant time calculations for complete weeks and remaining days

- **Space Complexity**: O(1)
  - Only a few variables needed

**Simulation Approach:**
- **Time Complexity**: O(n)
  - Iterate through n days

- **Space Complexity**: O(1)
  - Only tracking variables needed

### Implementation

```go
func totalMoney(n int) int {
    
}
```

### Test Cases

- `n = 4` → `10`
  - Day 1-4: 1 + 2 + 3 + 4 = 10

- `n = 10` → `37`
  - Week 1: 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28
  - Week 2 (days 1-3): 2 + 3 + 4 = 9
  - Total: 28 + 9 = 37

- `n = 20` → `96`
  - Week 1: 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28
  - Week 2: 2 + 3 + 4 + 5 + 6 + 7 + 8 = 35
  - Week 3 (days 1-6): 3 + 4 + 5 + 6 + 7 + 8 = 33
  - Total: 28 + 35 + 33 = 96

