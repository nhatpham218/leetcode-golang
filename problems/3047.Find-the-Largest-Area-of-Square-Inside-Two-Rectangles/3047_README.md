# 3047. Find the Largest Area of Square Inside Two Rectangles

## Problem Description

You are given `n` axis-aligned rectangles on the 2D plane, represented by two integer arrays `bottomLeft` and `topRight`, each of size `n x 2`. For the `i`-th rectangle:

- `bottomLeft[i] = [x1, y1]` is the bottom-left corner
- `topRight[i] = [x2, y2]` is the top-right corner

You may select any region formed by the intersection of **two** of these rectangles (not necessarily disjoint). Find the **largest square** (by area) that can be inscribed in such an intersecting region.

Return the **largest** such square area. If no two rectangles overlap, return `0`.

## Examples

### Example 1:
```
Input: bottomLeft = [[1,1],[2,2],[3,1]], topRight = [[3,3],[4,4],[6,6]]
Output: 1
Explanation: A square with side length 1 can be placed inside the intersection of rectangles 0 and 1, or inside the intersection of rectangles 1 and 2. Hence the largest area is 1 * 1 = 1.
```

### Example 2:
```
Input: bottomLeft = [[1,1],[2,2],[1,2]], topRight = [[3,3],[4,4],[3,4]]
Output: 1
Explanation: A square with side length 1 can be placed inside the intersection of rectangles 0 and 1, or inside the intersection of rectangles 0 and 2, or inside the intersection of rectangles 1 and 2. Hence the largest area is 1 * 1 = 1.
```

### Example 3:
```
Input: bottomLeft = [[1,1],[3,3],[3,1]], topRight = [[2,2],[4,4],[4,2]]
Output: 0
Explanation: No pair of rectangles intersect, hence return 0.
```

## Constraints

- `n == bottomLeft.length == topRight.length`
- `2 <= n <= 10^3`
- `bottomLeft[i].length == topRight[i].length == 2`
- `1 <= bottomLeft[i][0], bottomLeft[i][1] <= 10^7`
- `1 <= topRight[i][0], topRight[i][1] <= 10^7`
- `bottomLeft[i][0] < topRight[i][0]`
- `bottomLeft[i][1] < topRight[i][1]`

## Solution

To be implemented.
