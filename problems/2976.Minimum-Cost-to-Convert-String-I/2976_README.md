# 2976. Minimum Cost to Convert String I

## Problem Description

You are given two **0-indexed** strings `source` and `target`, both of length `n` and consisting of **lowercase** English letters. You are also given two **0-indexed** character arrays `original` and `changed`, and an integer array `cost`, where `cost[i]` represents the cost of changing the character `original[i]` to the character `changed[i]`.

You start with the string `source`. In one operation, you can pick a character `x` from the string and change it to the character `y` at a cost of `z` **if** there exists **any** index `j` such that `cost[j] == z`, `original[j] == x`, and `changed[j] == y`.

Return *the **minimum** cost to convert the string* `source` *to the string* `target` *using **any** number of operations. If it is impossible to convert* `source` *to* `target`, *return* `-1`.

**Note** that there may exist indices `i`, `j` such that `original[j] == original[i]` and `changed[j] == changed[i]`.

## Examples

### Example 1:
```
Input: source = "abcd", target = "acbe", original = ["a","b","c","c","e","d"], changed = ["b","c","b","e","b","e"], cost = [2,5,5,1,2,20]
Output: 28
Explanation: To convert the string "abcd" to string "acbe":
- Change value at index 1 from 'b' to 'c' at a cost of 5.
- Change value at index 2 from 'c' to 'e' at a cost of 1.
- Change value at index 2 from 'e' to 'b' at a cost of 2.
- Change value at index 3 from 'd' to 'e' at a cost of 20.
The total cost incurred is 5 + 1 + 2 + 20 = 28.
It can be shown that this is the minimum possible cost.
```

### Example 2:
```
Input: source = "aaaa", target = "bbbb", original = ["a","c"], changed = ["c","b"], cost = [1,2]
Output: 12
Explanation: To change the character 'a' to 'b' change the character 'a' to 'c' at a cost of 1, followed by changing the character 'c' to 'b' at a cost of 2, for a total cost of 1 + 2 = 3. To change all occurrences of 'a' to 'b', a total cost of 3 * 4 = 12 is incurred.
```

### Example 3:
```
Input: source = "abcd", target = "abce", original = ["a"], changed = ["e"], cost = [10000]
Output: -1
Explanation: It is impossible to convert source to target because the value at index 3 cannot be changed from 'd' to 'e'.
```

## Constraints

- `1 <= source.length == target.length <= 10^5`
- `source`, `target` consist of lowercase English letters.
- `1 <= cost.length == original.length == changed.length <= 2000`
- `original[i]`, `changed[i]` are lowercase English letters.
- `1 <= cost[i] <= 10^6`
- `original[i] != changed[i]`

## Hints

1. Construct a graph with each letter as a node, and construct an edge `(a, b)` with weight `c` if we can change from character `a` to letter `b` with cost `c`. (Keep the one with the smallest cost in case there are multiple edges between `a` and `b`).
2. Calculate the shortest path for each pair of characters `(source[i], target[i])`. The sum of cost over all `i` in the range `[0, source.length - 1]`. If there is no path between `source[i]` and `target[i]`, the answer is `-1`.
3. Any shortest path algorithms will work since we only have `26` nodes. Since we only have at most `26 * 26` pairs, we can save the result to avoid re-calculation.
4. We can also use Floyd Warshall's algorithm to precompute all the results.

## Solution

### Algorithm Overview

This problem can be modeled as a **shortest path problem** on a graph of 26 lowercase letters.

### Approach

1. **Build Graph**: Create a directed weighted graph where each node is a letter (a-z), and each edge represents a character transformation with its cost.
2. **Floyd-Warshall**: Since we only have 26 nodes, use Floyd-Warshall algorithm to precompute all-pairs shortest paths in O(26³) = O(1) time.
3. **Sum Costs**: For each position i, add the shortest path cost from `source[i]` to `target[i]`. If any path doesn't exist, return -1.

### Complexity Analysis

- **Time Complexity**: O(26³ + n) ≈ O(n) where n is the length of source/target
- **Space Complexity**: O(26²) = O(1) for the distance matrix

### Implementation

```go
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
    const inf = int64(1e18)

    // Initialize distance matrix for 26 letters
    dist := [26][26]int64{}
    for i := 0; i < 26; i++ {
        for j := 0; j < 26; j++ {
            if i == j {
                dist[i][j] = 0
            } else {
                dist[i][j] = inf
            }
        }
    }

    // Build edges from the transformation rules (keep minimum cost)
    for i := 0; i < len(original); i++ {
        from := int(original[i] - 'a')
        to := int(changed[i] - 'a')
        if int64(cost[i]) < dist[from][to] {
            dist[from][to] = int64(cost[i])
        }
    }

    // Floyd-Warshall: compute all-pairs shortest paths
    for k := 0; k < 26; k++ {
        for i := 0; i < 26; i++ {
            for j := 0; j < 26; j++ {
                if dist[i][k]+dist[k][j] < dist[i][j] {
                    dist[i][j] = dist[i][k] + dist[k][j]
                }
            }
        }
    }

    // Calculate total cost
    var totalCost int64
    for i := 0; i < len(source); i++ {
        from := int(source[i] - 'a')
        to := int(target[i] - 'a')
        if dist[from][to] == inf {
            return -1
        }
        totalCost += dist[from][to]
    }

    return totalCost
}
```

### Test Results

All test cases pass successfully:
- Example 1: ✅ Expected 28, Got 28
- Example 2: ✅ Expected 12, Got 12
- Example 3: ✅ Expected -1, Got -1
