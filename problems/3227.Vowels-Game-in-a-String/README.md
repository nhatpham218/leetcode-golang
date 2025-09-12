# 3227. Vowels Game in a String

## Problem Description

Alice and Bob are playing a game on a string.

You are given a string `s`, Alice and Bob will take turns playing the following game where Alice starts first:

- On Alice's turn, she has to remove any non-empty substring from `s` that contains an **odd** number of vowels.
- On Bob's turn, he has to remove any non-empty substring from `s` that contains an **even** number of vowels.

The first player who cannot make a move on their turn loses the game. We assume that both Alice and Bob play optimally.

Return `true` if Alice wins the game, and `false` otherwise.

The English vowels are: `a`, `e`, `i`, `o`, and `u`.

## Examples

### Example 1:
```
Input: s = "leetcoder"
Output: true
Explanation:
Alice can win the game as follows:
- Alice plays first, she can delete the underlined substring in s = "leetcoder" which contains 3 vowels. The resulting string is s = "der".
- Bob plays second, he can delete the underlined substring in s = "der" which contains 0 vowels. The resulting string is s = "er".
- Alice plays third, she can delete the whole string s = "er" which contains 1 vowel.
- Bob plays fourth, since the string is empty, there is no valid play for Bob. So Alice wins the game.
```

### Example 2:
```
Input: s = "bbcd"
Output: false
Explanation:
There is no valid play for Alice in her first turn, so Alice loses the game.
```

## Constraints

- `1 <= s.length <= 10^5`
- `s` consists only of lowercase English letters.

## Solution

### Algorithm Overview

This is a **Game Theory** problem that can be solved by analyzing the game's structure.

### Key Insights

1. **Alice needs odd vowels, Bob needs even vowels**
2. **Game analysis**: The key insight is that if there are any vowels in the string, Alice can always make a move by selecting a single vowel character (which gives her 1 vowel = odd number)
3. **Winning condition**: Alice wins if and only if there is at least one vowel in the string

### Approach

The solution is surprisingly simple:
- Count the total number of vowels in the string
- If there are any vowels (count > 0), Alice wins
- If there are no vowels, Alice cannot make any move and loses

### Why this works:

1. **If no vowels exist**: Alice cannot make any move (needs odd vowels), so she loses immediately
2. **If vowels exist**: Alice can always pick a single vowel character, guaranteeing she can make at least one move
3. **Optimal play**: With optimal play, the existence of vowels determines the winner

### Complexity Analysis

- **Time Complexity**: O(n) - Single pass through the string to count vowels
- **Space Complexity**: O(1) - Only using constant extra space

### Implementation

```go
func doesAliceWin(s string) bool {
    // Count vowels in the string
    // Return true if any vowels exist, false otherwise
}
```

### Test Cases

- `"leetcoder"` → `true` (contains vowels: e,e,o,e)
- `"bbcd"` → `false` (no vowels)
- `"a"` → `true` (single vowel)
- `"b"` → `false` (single consonant)
- `"aeiou"` → `true` (all vowels)
- `"bcdfg"` → `false` (all consonants)
