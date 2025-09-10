# 1733. Minimum Number of People to Teach

## Problem Description

On a social network consisting of m users and some friendships between users, two users can communicate with each other if they know a common language.

You are given an integer `n`, an array `languages`, and an array `friendships` where:

- There are `n` languages numbered 1 through n,
- `languages[i]` is the set of languages the ith user knows, and
- `friendships[i] = [ui, vi]` denotes a friendship between the users ui and vi.

You can choose one language and teach it to some users so that all friends can communicate with each other. Return the minimum number of users you need to teach.

Note that friendships are not transitive, meaning if x is a friend of y and y is a friend of z, this doesn't guarantee that x is a friend of z.

## Examples

### Example 1:
```
Input: n = 2, languages = [[1],[2],[1,2]], friendships = [[1,2],[1,3],[2,3]]
Output: 1
Explanation: You can either teach user 1 the second language or user 2 the first language.
```

### Example 2:
```
Input: n = 3, languages = [[2],[1,3],[1,2],[3]], friendships = [[1,4],[1,2],[3,4],[2,3]]
Output: 2
Explanation: Teach the third language to users 1 and 3, yielding two users to teach.
```

## Constraints

- `2 <= n <= 500`
- `languages.length == m`
- `1 <= m <= 500`
- `1 <= languages[i].length <= n`
- `1 <= languages[i][j] <= n`
- `1 <= ui < vi <= languages.length`
- `1 <= friendships.length <= 500`
- All tuples `(ui, vi)` are unique
- `languages[i]` contains only unique values

## Solution

### Algorithm Overview

This problem is solved using a **Greedy Approach** with **Set Operations**. The key insight is to find all problematic friendship pairs (those who cannot communicate) and then determine which language, when taught to the minimum number of users, resolves all communication issues.

### Approach

1. **Identify Communication Issues**: Find all friendship pairs that cannot communicate (no common language)
2. **Try Each Language**: For each possible language, calculate how many users need to learn it
3. **Minimize Teaching**: Choose the language that requires teaching the fewest users
4. **Handle Edge Cases**: Return 0 if all friends can already communicate

### Key Components

#### Language Set Conversion
```go
langSet := make([]map[int]bool, m)
for i := 0; i < m; i++ {
    langSet[i] = make(map[int]bool)
    for _, lang := range languages[i] {
        langSet[i][lang] = true
    }
}
```

#### Communication Check
```go
canCommunicate := func(u, v int) bool {
    for lang := range langSet[u] {
        if langSet[v][lang] {
            return true
        }
    }
    return false
}
```

#### Optimal Language Selection
For each language, determine users who need to learn it:
- If neither user in a problematic pair knows the language: both must learn it
- If only one user knows the language: the other must learn it
- If both know the language: no additional teaching needed

### Complexity Analysis

- **Time Complexity**: O(n × m × L + F × L)
  - Converting languages to sets: O(m × L) where L is avg languages per user
  - Finding problematic pairs: O(F × L) where F is number of friendships
  - Trying each language: O(n × P) where P is problematic pairs
- **Space Complexity**: O(m × L) for language sets and O(P) for problematic pairs

### Implementation Strategy

1. **Preprocessing**: Convert language arrays to hash sets for O(1) lookup
2. **Problem Identification**: Find all friendship pairs that cannot communicate
3. **Language Evaluation**: For each possible language, count users who need to learn it
4. **Optimization**: Return the minimum count across all languages

### Test Results

All test cases pass successfully:
- Example 1: ✅ Expected 1, Got 1
- Example 2: ✅ Expected 2, Got 2
- All can communicate: ✅ Expected 0, Got 0
- Single language optimal: ✅ Expected 2, Got 2
- Complex case: ✅ Expected 2, Got 2
- Edge case - two users: ✅ Expected 1, Got 1
