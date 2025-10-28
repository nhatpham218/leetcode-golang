# 3354. Make Array Elements Equal to Zero

## Problem Description

You are given an integer array nums.

Start by selecting a starting position curr such that nums[curr] == 0, and choose a movement direction of either left or right.

After that, you repeat the following process:

- If curr is out of the range [0, n - 1], this process ends.
- If nums[curr] == 0, move in the current direction by incrementing curr if you are moving right, or decrementing curr if you are moving left.
- Else if nums[curr] > 0:
  - Decrement nums[curr] by 1.
  - Reverse your movement direction (left becomes right and vice versa).
  - Take a step in your new direction.

A selection of the initial position curr and movement direction is considered valid if every element in nums becomes 0 by the end of the process.

Return the number of possible valid selections.

## Examples

### Example 1:
```
Input: nums = [1,0,2,0,3]
Output: 2
Explanation:
The only possible valid selections are the following:
- Choose curr = 3, and a movement direction to the left.
  [1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,1,0,3] -> [1,0,1,0,3] -> [1,0,1,0,2] -> [1,0,1,0,2] -> [1,0,0,0,2] -> [1,0,0,0,2] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,0].
- Choose curr = 3, and a movement direction to the right.
  [1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,2,0,2] -> [1,0,2,0,2] -> [1,0,1,0,2] -> [1,0,1,0,2] -> [1,0,1,0,1] -> [1,0,1,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [0,0,0,0,0].
```

### Example 2:
```
Input: nums = [2,3,4,0,4,1,0]
Output: 0
Explanation:
There are no possible valid selections.
```

## Constraints

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 100`
- There is at least one element i where nums[i] == 0.

## Solution

### Algorithm Overview

This problem requires simulating the process of moving through the array and decrementing values based on specific rules to determine valid starting positions and directions.

### Approach

1. **Find Zero Positions**: Identify all positions where nums[i] == 0 as potential starting points.
2. **Try Both Directions**: For each zero position, simulate the process moving left and moving right.
3. **Simulate Movement**: Follow the rules to move through the array, decrement values, and reverse direction when needed.
4. **Validate**: Check if all elements become zero after the process completes.
5. **Count Valid**: Count the number of valid (position, direction) pairs.

### Key Points

- Only positions with value 0 can be starting points
- Must simulate the exact process for each starting position and direction
- Direction reverses when encountering a positive value
- Process ends when moving out of bounds
- Need to check if all elements are zero after the process

### Complexity Analysis

- **Time Complexity**: O(n * m)
  - n is the number of zero positions
  - m is the maximum number of operations needed for simulation

- **Space Complexity**: O(n)
  - Need to copy the array for each simulation

### Implementation

```go
func countValidSelections(nums []int) int {
	
}
```

### Test Cases

- Array with multiple zeros and valid selections
- Array with no valid selections
- Array with only zeros
- Array with single element
- Edge cases with direction reversals

