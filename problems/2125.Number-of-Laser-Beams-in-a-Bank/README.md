# 2125. Number of Laser Beams in a Bank

## Problem Description

Anti-theft security devices are activated inside a bank. You are given a 0-indexed binary string array bank representing the floor plan of the bank, which is an m x n 2D matrix. bank[i] represents the ith row, consisting of '0's and '1's. '0' means the cell is empty, while '1' means the cell has a security device.

There is one laser beam between any two security devices if both conditions are met:

- The two devices are located on two different rows: r1 and r2, where r1 < r2.
- For each row i where r1 < i < r2, there are no security devices in the ith row.

Laser beams are independent, i.e., one beam does not interfere nor join with another.

Return the total number of laser beams in the bank.

## Examples

### Example 1:
```
Input: bank = ["011001","000000","010100","001000"]
Output: 8
Explanation: Between each of the following device pairs, there is one beam. In total, there are 8 beams:
 * bank[0][1] -- bank[2][1]
 * bank[0][1] -- bank[2][3]
 * bank[0][2] -- bank[2][1]
 * bank[0][2] -- bank[2][3]
 * bank[0][5] -- bank[2][1]
 * bank[0][5] -- bank[2][3]
 * bank[2][1] -- bank[3][2]
 * bank[2][3] -- bank[3][2]
Note that there is no beam between any device on the 0th row with any on the 3rd row.
This is because the 2nd row contains security devices, which breaks the second condition.
```

### Example 2:
```
Input: bank = ["000","111","000"]
Output: 0
Explanation: There does not exist two devices located on two different rows.
```

## Constraints

- `m == bank.length`
- `n == bank[i].length`
- `1 <= m, n <= 500`
- `bank[i][j]` is either '0' or '1'.

## Solution

### Algorithm Overview

This problem requires counting laser beams between security devices on different rows where there are no devices in the rows between them.

### Approach

1. **Count Devices**: For each row, count the number of security devices ('1's).
2. **Calculate Beams**: The number of beams between two consecutive rows with devices is the product of their device counts.
3. **Skip Empty Rows**: Only consider rows that have at least one device.
4. **Accumulate Total**: Sum up all the beams between consecutive non-empty rows.

### Key Points

- Laser beams only form between two rows with security devices
- Rows with no devices don't affect the beam count
- The number of beams between two rows is the product of their device counts
- Process rows sequentially and only track the previous row's device count

### Complexity Analysis

- **Time Complexity**: O(m * n)
  - Iterate through each row and count devices in each row

- **Space Complexity**: O(1)
  - Only need to track the previous row's device count

### Implementation

```go
func numberOfBeams(bank []string) int {
	
}
```

### Test Cases

- Multiple rows with devices separated by empty rows
- Consecutive rows with devices
- All rows empty
- Single row with devices
- Multiple devices in rows

