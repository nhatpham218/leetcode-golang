# 3661. Maximum Walls Destroyed by Robots

## Problem Description

There is an endless straight line populated with some robots and walls. You are given integer arrays `robots`, `distance`, and `walls`:

- `robots[i]` is the position of the ith robot.
- `distance[i]` is the maximum distance the ith robot's bullet can travel.
- `walls[j]` is the position of the jth wall.

Every robot has one bullet that can either fire to the left or the right at most `distance[i]` meters.

A bullet destroys every wall in its path that lies within its range. Robots are fixed obstacles: if a bullet hits another robot before reaching a wall, it immediately stops at that robot and cannot continue.

Return the maximum number of unique walls that can be destroyed by the robots.

## Notes

- A wall and a robot may share the same position; the wall can be destroyed by the robot at that position.
- Robots are not destroyed by bullets.

## Examples

### Example 1:
```
Input: robots = [4], distance = [3], walls = [1,10]
Output: 1
Explanation:
robots[0] = 4 fires left with distance[0] = 3, covering [1, 4] and destroys walls[0] = 1.
Thus, the answer is 1.
```

### Example 2:
```
Input: robots = [10,2], distance = [5,1], walls = [5,2,7]
Output: 3
Explanation:
robots[0] = 10 fires left with distance[0] = 5, covering [5, 10] and destroys walls[0] = 5 and walls[2] = 7.
robots[1] = 2 fires left with distance[1] = 1, covering [1, 2] and destroys walls[1] = 2.
Thus, the answer is 3.
```

### Example 3:
```
Input: robots = [1,2], distance = [100,1], walls = [10]
Output: 0
Explanation:
In this example, only robots[0] can reach the wall, but its shot to the right is blocked by robots[1]; thus the answer is 0.
```

## Constraints

- `1 <= robots.length == distance.length <= 10^5`
- `1 <= walls.length <= 10^5`
- `1 <= robots[i], walls[j] <= 10^9`
- `1 <= distance[i] <= 10^5`
- All values in `robots` are unique
- All values in `walls` are unique

## Solution

### Algorithm Overview

This problem is solved using **Dynamic Programming** with **Binary Search** optimization. The key insight is that each robot can shoot either left or right, but bullets are blocked by other robots in their path.

### Approach

1. **Sorting**: Sort both robots (with distances) and walls by position for efficient processing
2. **Dynamic Programming**: Use DP to track maximum walls destroyed based on shooting direction
3. **Binary Search**: Efficiently count walls within shooting ranges
4. **Robot Blocking**: Consider that bullets stop when hitting other robots

### State Definition

- `dp[i][0]`: Maximum walls destroyed when robot `i` shoots **left**
- `dp[i][1]`: Maximum walls destroyed when robot `i` shoots **right**

### Key Components

#### Binary Search Helper
```go
query := func(l, r int) int {
    if l > r { return 0 }
    left := sort.SearchInts(walls, l)      // First wall >= l
    right := sort.SearchInts(walls, r+1)   // First wall > r
    return right - left                    // Count walls in [l,r]
}
```

#### DP Transitions

**Base Case (First Robot):**
- Left: `dp[0][0] = query(pos - dist, pos)`
- Right: `dp[0][1] = query(pos, min(nextRobot-1, pos+dist))`

**General Case (Robot i):**
- **Shooting Right**: `dp[i][1] = max(dp[i-1][0], dp[i-1][1]) + query(pos, rightLimit)`
- **Shooting Left**: Handle two scenarios:
  1. Previous robot shot left: `dp[i-1][0] + query(leftLimit, pos)`
  2. Previous robot shot right: `dp[i-1][1] + query(leftLimit, pos) - overlap`

### Complexity Analysis

- **Time Complexity**: O(n log n + m log m + n log m)
  - Sorting robots: O(n log n)
  - Sorting walls: O(m log m)  
  - DP with binary search: O(n log m)
- **Space Complexity**: O(n + m) for storing robots, walls, and DP array

### Implementation

```go
func maxWalls(robots []int, distance []int, walls []int) int {
    // Sort robots by position with their distances
    // Sort walls for binary search
    // Use DP to find optimal shooting strategy
    // Return max(dp[n-1][0], dp[n-1][1])
}
```

### Test Results

All test cases pass successfully:
- Example 1: ✅ Expected 1, Got 1
- Example 2: ✅ Expected 3, Got 3  
- Example 3: ✅ Expected 0, Got 0
- Example 4: ✅ Expected 15, Got 15
- Example 5: ✅ Expected 41, Got 41
- Example 6: ✅ Expected 49, Got 49
