# 3347. Maximum Frequency of an Element After Performing Operations II

## Problem Description

You are given an integer array `nums` and two integers `k` and `numOperations`.

You must perform an operation `numOperations` times on `nums`, where in each operation you:

- Select an index `i` that was not selected in any previous operations.
- Add an integer in the range `[-k, k]` to `nums[i]`.

Return the maximum possible frequency of any element in `nums` after performing the operations.

## Examples

### Example 1:
```
Input: nums = [1,4,5], k = 1, numOperations = 2
Output: 2
Explanation:
We can achieve a maximum frequency of two by:
- Adding 0 to nums[1], after which nums becomes [1, 4, 5].
- Adding -1 to nums[2], after which nums becomes [1, 4, 4].
```

### Example 2:
```
Input: nums = [5,11,20,20], k = 5, numOperations = 1
Output: 2
Explanation:
We can achieve a maximum frequency of two by:
- Adding 0 to nums[1].
```

## Constraints

- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^9`
- `0 <= k <= 10^9`
- `0 <= numOperations <= nums.length`

## Solution

### Algorithm Overview

The problem requires finding the maximum frequency of any element after performing at most `numOperations` operations, where each operation can add a value in the range `[-k, k]` to an element.

The solution uses a **two-phase sliding window approach** that considers two types of potential target values:
1. **Existing values** in the array (values that already have some frequency)
2. **Virtual targets** (values not in array where elements can converge)

### Key Insights

1. **Transformation Range**: An element `nums[i]` can be transformed to any value in the range `[nums[i] - k, nums[i] + k]`.

2. **Overlapping Intervals**: If two elements are within `2*k` distance of each other, they can potentially converge to a common value between them.

3. **Greedy Approach**: For any target value, we greedily use operations on elements closest to the target (within range `[target-k, target+k]`).

4. **Monotonic Pointers**: Since the array is sorted, we can use sliding window technique with left and right pointers that only move forward.

### Approach

#### Phase 1: Existing Values as Targets

For each unique value `x` in the sorted array:
- Calculate its current frequency `freq`
- Define the range `[x-k, x+k]` of elements that can be converted to `x`
- Use sliding window to count total elements in this range: `totalInRange`
- Maximum frequency achievable: `freq + min(totalInRange - freq, numOperations)`
  - `freq`: elements already at target
  - `totalInRange - freq`: elements needing conversion
  - Can convert at most `numOperations` elements

**Optimization**: Process each unique value once by skipping duplicates.

#### Phase 2: Virtual Targets (Overlapping Intervals)

For each position in the sorted array, consider a sliding window of size `2*k`:
- Elements in range `[currVal - 2*k, currVal]` can all converge to some value
- If we have `n` elements in this range, they can meet at a point in the middle
- Maximum frequency: `min(window_size, numOperations)`
- This captures scenarios where no single existing value is optimal

**Example**: `nums = [5, 64], k = 42`
- Element 5 can reach `[5-42, 5+42]` = `[-37, 47]`
- Element 64 can reach `[64-42, 64+42]` = `[22, 106]`
- Overlap = `[22, 47]` (any value here can be reached by both)
- Phase 2 detects this: when processing 64, window `[64-84, 64]` = `[-20, 64]` contains both elements
- Result: 2 operations → frequency of 2

### Complexity Analysis

- **Time Complexity**: O(n log n)
  - Sorting: O(n log n)
  - Phase 1: O(n) - each pointer moves at most n times
  - Phase 2: O(n) - single pass with sliding window
  - Total: O(n log n)

- **Space Complexity**: O(1)
  - Only constant extra space for pointers and counters
  - Sorting is in-place

### Implementation

```go
func maxFrequency(nums []int, k int, numOperations int) int {
	n := len(nums)
	sort.Ints(nums)

	maxFreq := 0
	left := 0
	right := 0

	// Phase 1: Consider existing values as targets
	for i := 0; i < n; i++ {
		currVal := nums[i]
		leftBound := max(1, currVal-k)
		rightBound := min(nums[n-1], currVal+k)

		// Count frequency of current value
		freq := 1
		nextIdx := i + 1
		for nextIdx < n && nums[nextIdx] == currVal {
			freq++
			nextIdx++
		}
		i = nextIdx - 1

		// Move left pointer to first element >= leftBound
		for left < n && nums[left] < leftBound {
			left++
		}

		// Move right pointer to last element <= rightBound
		right = max(right, i)
		for right+1 < n && nums[right+1] <= rightBound {
			right++
		}

		// Calculate max frequency for this target
		totalInRange := right - left + 1
		maxFreq = max(maxFreq, freq+min(totalInRange-freq, numOperations))
	}

	// Phase 2: Consider virtual targets (overlapping intervals)
	for left, right = 0, 0; right < n; right++ {
		currVal := nums[right]
		leftBound := max(1, currVal-2*k)
		for left < right && nums[left] < leftBound {
			left++
		}
		maxFreq = max(maxFreq, min(right-left+1, numOperations))
	}

	return maxFreq
}
```

### Test Cases

- `nums = [1,4,5], k = 1, numOperations = 2` → `2`
  - Target 4: elements 4 and 5 can become 4 → frequency 2
  
- `nums = [5,11,20,20], k = 5, numOperations = 1` → `2`
  - Target 20: already has frequency 2
  
- `nums = [5,64], k = 42, numOperations = 2` → `2`
  - Virtual target ~47: both elements can reach it → frequency 2


