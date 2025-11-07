# 2528. Maximize the Minimum Powered City

## Problem Description

You are given a 0-indexed integer array `stations` of length `n`, where `stations[i]` represents the number of power stations in the `ith` city.

Each power station can provide power to every city in a fixed range. In other words, if the range is denoted by `r`, then a power station at city `i` can provide power to all cities `j` such that `|i - j| <= r` and `0 <= i, j <= n - 1`.

**Note** that `|x|` denotes absolute value. For example, `|7 - 5| = 2` and `|3 - 10| = 7`.

The **power** of a city is the total number of power stations it is being provided power from.

The government has sanctioned building `k` more power stations, each of which can be built in any city, and have the same range as the pre-existing ones.

Given the two integers `r` and `k`, return the **maximum possible minimum power** of a city, if the additional power stations are built optimally.

**Note** that you can build the `k` power stations in multiple cities.

## Examples

### Example 1:
```
Input: stations = [1,2,4,5,0], r = 1, k = 2
Output: 5
Explanation: 
One of the optimal ways is to install both the power stations at city 1. 
So stations will become [1,4,4,5,0].
- City 0 is provided by 1 + 4 = 5 power stations.
- City 1 is provided by 1 + 4 + 4 = 9 power stations.
- City 2 is provided by 4 + 4 + 5 = 13 power stations.
- City 3 is provided by 5 + 4 = 9 power stations.
- City 4 is provided by 5 + 0 = 5 power stations.
So the minimum power of a city is 5.
Since it is not possible to obtain a larger power, we return 5.
```

### Example 2:
```
Input: stations = [4,4,4,4], r = 0, k = 3
Output: 4
Explanation: 
It can be proved that we cannot make the minimum power of a city greater than 4.
```

## Constraints

- `n == stations.length`
- `1 <= n <= 10^5`
- `0 <= stations[i] <= 10^5`
- `0 <= r <= n - 1`
- `0 <= k <= 10^9`

## Solution

### Algorithm Overview

This problem requires maximizing the minimum power across all cities by optimally placing k additional power stations.

### Key Insights

1. **Binary Search**: The answer can be binary searched. If we can achieve a minimum power of `mid`, we can also achieve any value less than `mid`.
2. **Greedy Placement**: For a given target minimum power, we can greedily place power stations to achieve it.
3. **Range Coverage**: Each power station at position `i` covers cities from `max(0, i-r)` to `min(n-1, i+r)`.

### Approach

1. Binary search on the answer (minimum power)
2. For each candidate answer:
   - Check if it's achievable with k power stations
   - Use greedy strategy: place stations as far right as possible to maximize coverage
3. Return the maximum achievable minimum power

### Complexity Analysis

- **Time Complexity**: O(n * log(sum(stations) + k))
- **Space Complexity**: O(n)

### Implementation

```go
func maxPower(stations []int, r int, k int) int64 {
    // Binary search on answer
}
```

### Test Cases

- `stations = [1,2,4,5,0], r = 1, k = 2` → `5`
- `stations = [4,4,4,4], r = 0, k = 3` → `4`



