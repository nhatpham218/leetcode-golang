# 2011. Final Value of Variable After Performing Operations

## Problem Description

There is a programming language with only four operations and one variable X:

- `++X` and `X++` increments the value of the variable X by 1.
- `--X` and `X--` decrements the value of the variable X by 1.

Initially, the value of X is 0.

Given an array of strings `operations` containing a list of operations, return the final value of X after performing all the operations.

## Examples

### Example 1:
```
Input: operations = ["--X","X++","X++"]
Output: 1
Explanation: The operations are performed as follows:
Initially, X = 0.
--X: X is decremented by 1, X =  0 - 1 = -1.
X++: X is incremented by 1, X = -1 + 1 =  0.
X++: X is incremented by 1, X =  0 + 1 =  1.
```

### Example 2:
```
Input: operations = ["++X","++X","X++"]
Output: 3
Explanation: The operations are performed as follows:
Initially, X = 0.
++X: X is incremented by 1, X = 0 + 1 = 1.
++X: X is incremented by 1, X = 1 + 1 = 2.
X++: X is incremented by 1, X = 2 + 1 = 3.
```

### Example 3:
```
Input: operations = ["X++","++X","--X","X--"]
Output: 0
Explanation: The operations are performed as follows:
Initially, X = 0.
X++: X is incremented by 1, X = 0 + 1 = 1.
++X: X is incremented by 1, X = 1 + 1 = 2.
--X: X is decremented by 1, X = 2 - 1 = 1.
X--: X is decremented by 1, X = 1 - 1 = 0.
```

## Constraints

- `1 <= operations.length <= 100`
- `operations[i]` will be either `"++X"`, `"X++"`, `"--X"`, or `"X--"`.

## Solution

### Algorithm Overview

The problem requires simulating the operations on a variable X that starts at 0.

### Key Insights

1. **Operations**: There are only 4 types of operations: `++X`, `X++`, `--X`, and `X--`.
2. **Effect**: `++X` and `X++` both increment X by 1, while `--X` and `X--` both decrement X by 1.
3. **Order**: The prefix/postfix notation doesn't matter for the final result.

### Approach

1. Initialize X to 0
2. For each operation, check if it contains `+` or `-`
3. Increment or decrement X accordingly
4. Return the final value of X

### Complexity Analysis

- **Time Complexity**: TBD
- **Space Complexity**: TBD

### Implementation

```go
func finalValueAfterOperations(operations []string) int {
    // Implementation to be added
}
```

### Test Cases

- `operations = ["--X","X++","X++"]` → `1`
- `operations = ["++X","++X","X++"]` → `3`
- `operations = ["X++","++X","--X","X--"]` → `0`

