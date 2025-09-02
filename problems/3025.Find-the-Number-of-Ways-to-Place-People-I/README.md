# 3025. Find the Number of Ways to Place People I

## Problem Description

You are given a 2D array `points` of size `n x 2` representing integer coordinates of some points on a 2D plane, where `points[i] = [xi, yi]`.

Count the number of pairs of points `(A, B)`, where:

- `A` is on the upper left side of `B`, and
- there are no other points in the rectangle (or line) they make (including the border).

Return the count.

## Examples

### Example 1:
```
Input: points = [[1,1],[2,2],[3,3]]
Output: 0

Explanation:
There is no way to choose A and B so A is on the upper left side of B.
```

### Example 2:
```
Input: points = [[6,2],[4,4],[2,6]]
Output: 2

Explanation:
The left one is the pair (points[1], points[0]), where points[1] is on the upper left side of points[0] and the rectangle is empty.
The middle one is the pair (points[2], points[1]), same as the left one it is a valid pair.
The right one is the pair (points[2], points[0]), where points[2] is on the upper left side of points[0], but points[1] is inside the rectangle so it's not a valid pair.
```

### Example 3:
```
Input: points = [[3,1],[1,3],[1,1]]
Output: 2

Explanation:
The left one is the pair (points[2], points[0]), where points[2] is on the upper left side of points[0] and there are no other points on the line they form. Note that it is a valid state when the two points form a line.
The middle one is the pair (points[1], points[2]), it is a valid pair same as the left one.
The right one is the pair (points[1], points[0]), it is not a valid pair as points[2] is on the border of the rectangle.
```

## Constraints

- `2 <= n <= 50`
- `points[i].length == 2`
- `0 <= points[i][0], points[i][1] <= 50`
- All `points[i]` are distinct.

## Solution

The solution function `numberOfPairs(points [][]int) int` is provided but left blank for you to implement.

### Algorithm Approach

1. For each pair of points (A, B), check if A is on the upper left side of B
2. A point A(x1, y1) is on the upper left side of point B(x2, y2) if `x1 <= x2` and `y1 >= y2`
3. For each valid pair, check if there are any other points inside or on the border of the rectangle formed by A and B
4. Count only the pairs where the rectangle/line is empty of other points

### Key Points

- A is upper left of B means: `A.x <= B.x` and `A.y >= B.y`
- Need to check all other points to ensure they don't lie within the rectangle (including borders)
- Rectangle boundaries are inclusive
- Points can form a line (when x coordinates or y coordinates are the same)
