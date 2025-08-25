# 2264. Largest 3-Same-Digit Number in String

## Problem Description

You are given a string `num` representing a large integer. An integer is **good** if it meets the following conditions:

- It is a substring of `num` with length 3.
- It consists of only one unique digit.

Return the maximum good integer as a string or an empty string `""` if no such integer exists.

**Note:**
- A substring is a contiguous sequence of characters within a string.
- There may be leading zeroes in `num` or a good integer.

## Examples

### Example 1:
```
Input: num = "6777133339"
Output: "777"
```
**Explanation:** The good integers are "777" and "333". "777" is the largest.

### Example 2:
```
Input: num = "2300019"
Output: "000"
```
**Explanation:** The only good integer is "000".

### Example 3:
```
Input: num = "42352338"
Output: ""
```
**Explanation:** No good integer exists.

## Solution Approach

The solution iterates through all possible 3-digit substrings of the input string and checks if they consist of the same digit. It keeps track of the largest good integer found and returns it at the end.

**Algorithm:**
1. Initialize an empty string to store the maximum good integer
2. Iterate through the string from index 0 to `len(num)-3`
3. For each position, extract the 3-character substring
4. Check if all three characters are the same
5. If it's a good integer, update the maximum if it's larger than the current maximum
6. Return the maximum good integer found

**Time Complexity:** O(n), where n is the length of the input string
**Space Complexity:** O(1), as we only store a constant amount of extra space

## Running the Solution

To run the solution:
```bash
go run solution.go
```

To run the tests:
```bash
go test
```
