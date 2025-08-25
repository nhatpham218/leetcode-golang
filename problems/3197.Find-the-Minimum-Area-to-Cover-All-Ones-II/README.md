# 3197. Find the Minimum Area to Cover All Ones II

## Problem Description

You are given a 2D binary array `grid`. You need to find 3 non-overlapping rectangles having non-zero areas with horizontal and vertical sides such that all the 1's in `grid` lie inside these rectangles.

Return the minimum possible sum of the area of these rectangles.

Note that the rectangles are allowed to touch.

## Examples

### Example 1:
```
Input: grid = [[1,0,1],[1,1,1]]
Output: 5
Explanation:
- The 1's at (0, 0) and (1, 0) are covered by a rectangle of area 2.
- The 1's at (0, 2) and (1, 2) are covered by a rectangle of area 2.
- The 1 at (1, 1) is covered by a rectangle of area 1.
```

### Example 2:
```
Input: grid = [[1,0,1,0],[0,1,0,1]]
Output: 5
Explanation:
- The 1's at (0, 0) and (0, 2) are covered by a rectangle of area 3.
- The 1 at (1, 1) is covered by a rectangle of area 1.
- The 1 at (1, 3) is covered by a rectangle of area 1.
```

## Constraints

- `1 <= grid.length, grid[i].length <= 30`
- `grid[i][j]` is either 0 or 1
- The input is generated such that there are at least three 1's in grid

## Solution Approach

The solution uses a brute force approach with bit manipulation to try all possible ways to partition the 1's into 3 groups:

1. **Find all positions of 1's**: First, we collect all coordinates where `grid[i][j] == 1`

2. **Try all possible partitions**: We use two masks to create 3 groups:
   - `mask1`: selects positions for the first group
   - `mask2`: selects positions for the second group (must be disjoint from mask1)
   - The remaining positions form the third group

3. **Ensure valid partitions**: Each group must have at least one position, and masks must be disjoint

4. **Calculate bounding areas**: For each group, calculate the minimum area rectangle needed to cover all positions in that group

5. **Find minimum total area**: Keep track of the minimum total area across all valid partitions

## Time Complexity

- **Time**: O(2^n × 2^n) = O(4^n) where n is the number of 1's
- **Space**: O(n) for storing positions and groups

## Key Insights

- The problem requires finding the optimal partition of 1's into exactly 3 groups
- Each group needs to be covered by a minimum bounding rectangle
- The total area is the sum of areas of all 3 rectangles
- We need to try all possible combinations to find the global minimum
- Using two masks ensures we create exactly 3 groups efficiently
