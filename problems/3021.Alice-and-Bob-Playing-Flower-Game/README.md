# 3021. Alice and Bob Playing Flower Game

## Problem Description

Alice and Bob are playing a turn-based game on a field, with two lanes of flowers between them. There are `x` flowers in the first lane between Alice and Bob, and `y` flowers in the second lane between them.

The game proceeds as follows:
- Alice takes the first turn.
- In each turn, a player must choose either one of the lane and pick one flower from that side.
- At the end of the turn, if there are no flowers left at all, the current player captures their opponent and wins the game.

Given two integers, `n` and `m`, the task is to compute the number of possible pairs `(x, y)` that satisfy the conditions:
- Alice must win the game according to the described rules.
- The number of flowers `x` in the first lane must be in the range `[1,n]`.
- The number of flowers `y` in the second lane must be in the range `[1,m]`.

Return the number of possible pairs `(x, y)` that satisfy the conditions mentioned in the statement.

## Examples

### Example 1:
```
Input: n = 3, m = 2
Output: 3
Explanation: The following pairs satisfy conditions described in the statement: (1,2), (3,2), (2,1).
```

### Example 2:
```
Input: n = 1, m = 1  
Output: 0
Explanation: No pairs satisfy the conditions described in the statement.
```

## Constraints

- `1 <= n, m <= 10^5`

## Algorithm

The key insight is that Alice wins if and only if the total number of flowers `x + y` is odd.

Since Alice goes first, she will make the last move (and win) when the total number of moves is odd.

For `x + y` to be odd, we need either:
1. `x` is odd and `y` is even
2. `x` is even and `y` is odd

**Counting:**
- Odd numbers in `[1,n]`: `(n + 1) / 2`
- Even numbers in `[1,n]`: `n / 2`
- Odd numbers in `[1,m]`: `(m + 1) / 2`  
- Even numbers in `[1,m]`: `m / 2`

**Answer:** `oddX × evenY + evenX × oddY`

**Time Complexity:** O(1)  
**Space Complexity:** O(1)
