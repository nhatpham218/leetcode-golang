# 1578. Minimum Time to Make Rope Colorful

## Problem Description

Alice has `n` balloons arranged on a rope. You are given a **0-indexed** string `colors` where `colors[i]` is the color of the `ith` balloon.

Alice wants the rope to be **colorful**. She does not want **two consecutive balloons** to be of the same color, so she asks Bob for help. Bob can remove some balloons from the rope to make it colorful. You are given a **0-indexed** integer array `neededTime` where `neededTime[i]` is the time (in seconds) that Bob needs to remove the `ith` balloon from the rope.

Return the **minimum time** Bob needs to make the rope colorful.

## Examples

### Example 1:
```
Input: colors = "abaac", neededTime = [1,2,3,4,5]
Output: 3
Explanation: In the above image, 'a' is blue, 'b' is red, and 'c' is green.
Bob can remove the blue balloon at index 2. This takes 3 seconds.
There are no longer two consecutive balloons of the same color. Total time = 3.
```

### Example 2:
```
Input: colors = "abc", neededTime = [1,2,3]
Output: 0
Explanation: The rope is already colorful. Bob does not need to remove any balloons from the rope.
```

### Example 3:
```
Input: colors = "aabaa", neededTime = [1,2,3,4,1]
Output: 2
Explanation: Bob will remove the balloons at indices 0 and 4. Each balloons takes 1 second to remove.
There are no longer two consecutive balloons of the same color. Total time = 1 + 1 = 2.
```

## Constraints

- `n == colors.length == neededTime.length`
- `1 <= n <= 10^5`
- `1 <= neededTime[i] <= 10^4`
- `colors` contains only lowercase English letters

## Solution

### Algorithm Overview

This problem requires finding the minimum cost to remove balloons such that no two consecutive balloons have the same color.

### Key Insights

1. **Greedy Approach**: When we encounter consecutive balloons of the same color, we need to remove all but one of them.
2. **Optimal Choice**: Among consecutive balloons of the same color, we should keep the one with the maximum `neededTime` and remove all others.
3. **Single Pass**: We can solve this problem in a single pass through the array.

### Approach

1. Iterate through the balloons
2. When we find consecutive balloons of the same color, identify the group
3. For each group, sum all times and subtract the maximum time (keep the most expensive balloon)
4. Add the cost to the total

### Complexity Analysis

- **Time Complexity**: O(n) where n is the number of balloons
- **Space Complexity**: O(1) constant space

### Implementation

```go
func minCost(colors string, neededTime []int) int {
    // Implement greedy approach
    // Keep the most expensive balloon in consecutive groups
}
```

### Test Cases

- `colors = "abaac", neededTime = [1,2,3,4,5]` → `3`
- `colors = "abc", neededTime = [1,2,3]` → `0`
- `colors = "aabaa", neededTime = [1,2,3,4,1]` → `2`

