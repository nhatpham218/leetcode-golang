# 3027. Find the Number of Ways to Place People II

## Problem Description

You are given a 2D array `points` of size `n x 2` representing integer coordinates of some points on a 2D-plane, where `points[i] = [xi, yi]`.

We define the right direction as positive x-axis (increasing x-coordinate) and the left direction as negative x-axis (decreasing x-coordinate). Similarly, we define the up direction as positive y-axis (increasing y-coordinate) and the down direction as negative y-axis (decreasing y-coordinate).

You have to place n people, including Alice and Bob, at these points such that there is exactly one person at every point. Alice wants to be alone with Bob, so Alice will build a rectangular fence with Alice's position as the upper left corner and Bob's position as the lower right corner of the fence (Note that the fence might not enclose any area, i.e. it can be a line). If any person other than Alice and Bob is either inside the fence or on the fence, Alice will be sad.

Return the number of pairs of points where you can place Alice and Bob, such that Alice does not become sad on building the fence.

Note that Alice can only build a fence with Alice's position as the upper left corner, and Bob's position as the lower right corner.

## Examples

### Example 1:
```
Input: points = [[1,1],[2,2],[3,3]]
Output: 0

Explanation: 
There is no way to place Alice and Bob such that Alice can build a fence with Alice's position as the upper left corner and Bob's position as the lower right corner. Hence we return 0.
```

### Example 2:
```
Input: points = [[6,2],[4,4],[2,6]]
Output: 2

Explanation: 
There are two ways to place Alice and Bob such that Alice will not be sad:
- Place Alice at (4, 4) and Bob at (6, 2).
- Place Alice at (2, 6) and Bob at (4, 4).
You cannot place Alice at (2, 6) and Bob at (6, 2) because the person at (4, 4) will be inside the fence.
```

### Example 3:
```
Input: points = [[3,1],[1,3],[1,1]]
Output: 2

Explanation:
There are two ways to place Alice and Bob such that Alice will not be sad:
- Place Alice at (1, 1) and Bob at (3, 1).
- Place Alice at (1, 3) and Bob at (1, 1).
You cannot place Alice at (1, 3) and Bob at (3, 1) because the person at (1, 1) will be on the fence.
```

## Constraints

- `2 <= n <= 1000`
- `points[i].length == 2`
- `-10^9 <= points[i][0], points[i][1] <= 10^9`
- All `points[i]` are distinct.

## Solution Approach

The key insight is that for Alice to be at the upper left corner and Bob at the lower right corner:
- Alice's x-coordinate ≤ Bob's x-coordinate (ax ≤ bx)
- Alice's y-coordinate ≥ Bob's y-coordinate (ay ≥ by)

### Algorithm:

1. **Sort the points**: Sort by x-coordinate ascending, then by y-coordinate descending. This ensures we process potential Alice positions from left to right, and for same x, from top to bottom.

2. **For each Alice position**: 
   - Consider it as the upper left corner
   - Track the maximum y-coordinate of valid Bob positions seen so far
   
3. **For each potential Bob position** (to the right of Alice):
   - Check if it can form a valid rectangle (ax ≤ bx and ay ≥ by)
   - Ensure no other point lies between Alice and Bob in the y-direction by checking if Bob's y > maxY
   - If valid, increment count and update maxY

### Time Complexity: O(n²)
### Space Complexity: O(1) excluding input

The sorting ensures we only need to track one variable (maxY) to determine if there are points inside the fence, making this an efficient solution for the given constraints.
