package leetcode

import (
	"fmt"
	"sort"
)

// 3646. Next Special Palindrome Number
// You are given an integer n.
// A number is called special if:
// - It is a palindrome.
// - Every digit k in the number appears exactly k times.
// Return the smallest special number strictly greater than n.

var specialNumbers []int64

func init() {
	// Generate all valid special palindrome numbers
	specialNumbers = generateAllValidSpecialNumbers()

	sort.Slice(specialNumbers, func(i, j int) bool {
		return specialNumbers[i] < specialNumbers[j]
	})
}

func generateAllValidSpecialNumbers() []int64 {
	var candidates []int64
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Use bitmasking to generate all possible combinations of digits 1-9
	for mask := 0; mask < (1 << 9); mask++ {
		var chosenElements []int

		// Check which digits are selected by the mask
		for pos := 0; pos < 9; pos++ {
			if mask&(1<<pos) != 0 {
				chosenElements = append(chosenElements, elements[pos])
			}
		}

		// Check validity constraints
		oddCount := 0
		totalLength := 0
		for _, val := range chosenElements {
			if val%2 == 1 {
				oddCount++
			}
			totalLength += val
		}

		// Skip if more than one odd digit (can't form valid palindrome)
		if oddCount > 1 {
			continue
		}

		// Skip if total length is too large (reasonable upper bound)
		if totalLength >= 17 {
			continue
		}

		// Generate palindromes for this combination
		palindromes := generatePalindromes(chosenElements)
		candidates = append(candidates, palindromes...)
	}

	// Remove duplicates and sort
	seen := make(map[int64]bool)
	var result []int64
	for _, num := range candidates {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func generatePalindromes(chosenElements []int) []int64 {
	var result []int64

	// Separate into half digits and odd digit
	var half []int
	oddElement := -1

	for _, val := range chosenElements {
		if val%2 == 1 {
			oddElement = val
		}
		count := val / 2
		for i := 0; i < count; i++ {
			half = append(half, val)
		}
	}

	// Generate all permutations of the half
	permutations := generatePermutations(half)

	for _, perm := range permutations {
		// Build palindrome: perm + oddElement + reverse(perm)
		var candidate int64 = 0

		// Add first half
		for _, val := range perm {
			candidate = candidate*10 + int64(val)
		}

		// Add odd element in middle if exists
		if oddElement != -1 {
			candidate = candidate*10 + int64(oddElement)
		}

		// Add second half (reverse of first half)
		for i := len(perm) - 1; i >= 0; i-- {
			candidate = candidate*10 + int64(perm[i])
		}

		result = append(result, candidate)
	}

	return result
}

func generatePermutations(arr []int) [][]int {
	if len(arr) == 0 {
		return [][]int{{}}
	}

	// Create a copy and sort to handle duplicates properly
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	var result [][]int
	var current []int
	used := make([]bool, len(sorted))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(sorted) {
			perm := make([]int, len(current))
			copy(perm, current)
			result = append(result, perm)
			return
		}

		for i := 0; i < len(sorted); i++ {
			if used[i] {
				continue
			}

			// Skip duplicates: if current element equals previous and previous wasn't used
			if i > 0 && sorted[i] == sorted[i-1] && !used[i-1] {
				continue
			}

			current = append(current, sorted[i])
			used[i] = true
			backtrack()
			current = current[:len(current)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}

func isValidSpecialNumber(num int64) bool {
	s := fmt.Sprintf("%d", num)

	// Check if it's a palindrome
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	// Count digit frequencies
	freq := make(map[int]int)
	for _, ch := range s {
		digit := int(ch - '0')
		freq[digit]++
	}

	// Check if each digit k appears exactly k times
	// For digit 0, it can only appear 0 times (which means it shouldn't appear at all)
	for digit, count := range freq {
		if digit == 0 {
			// Digit 0 cannot appear in a special number because 0 != count for any count > 0
			return false
		}
		if digit != count {
			return false
		}
	}

	return true
}

func specialPalindrome(n int64) int64 {
	// Binary search to find the smallest special number > n
	left, right := 0, len(specialNumbers)-1
	result := int64(-1)

	for left <= right {
		mid := (left + right) / 2
		if specialNumbers[mid] > n {
			result = specialNumbers[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}
