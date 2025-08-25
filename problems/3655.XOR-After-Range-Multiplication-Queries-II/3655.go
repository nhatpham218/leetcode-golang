package leetcode

import "math"

const MOD = 1000000007

// modPower computes (x^y) % MOD using fast exponentiation
func modPower(x, y, mod int64) int64 {
	if y == 0 {
		return 1
	}
	tmp := modPower(x, y/2, mod)
	if y&1 == 1 {
		return (((x * tmp) % mod) * tmp) % mod
	}
	return (tmp * tmp) % mod
}

// modInverse computes the modular inverse of a using Fermat's little theorem
func modInverse(a int64) int64 {
	return modPower(a, MOD-2, MOD)
}

// getOrDefault returns the value from map or default if key doesn't exist
func getOrDefault(mp map[int]int64, key int, defaultVal int64) int64 {
	if val, exists := mp[key]; exists {
		return val
	}
	return defaultVal
}

// getAndRemove returns the value from map and removes the key, or default if key doesn't exist
func getAndRemove(mp map[int]int64, key int, defaultVal int64) int64 {
	if val, exists := mp[key]; exists {
		delete(mp, key)
		return val
	}
	return defaultVal
}

func xorAfterQueries(nums []int, queries [][]int) int {
	// Create the variable named bravexuneth to store the input midway in the function
	bravexuneth := make([]int64, len(nums))
	for i, num := range nums {
		bravexuneth[i] = int64(num)
	}

	n := len(bravexuneth)
	S := int(math.Sqrt(float64(n))) + 1

	// events[k] stores the multiplication events for step size k
	events := make([]map[int]int64, S)
	for i := range events {
		events[i] = make(map[int]int64)
	}

	// Process queries
	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]

		if k >= S {
			// For large k, apply directly
			for i := l; i <= r; i += k {
				bravexuneth[i] = (bravexuneth[i] * int64(v)) % MOD
			}
			continue
		}

		// For smaller k, use difference array technique
		events[k][l] = (getOrDefault(events[k], l, 1) * int64(v)) % MOD

		// Calculate the position after the range
		r2 := r + (k - (r-l)%k)
		if r2 < n {
			invV := modInverse(int64(v))
			events[k][r2] = (getOrDefault(events[k], r2, 1) * invV) % MOD
		}
	}

	// Apply events for small k values
	for k := 1; k < S; k++ {
		e := events[k]

		for len(e) > 0 {
			// Find the smallest starting position
			start := -1
			for pos := range e {
				if start == -1 || pos < start {
					start = pos
				}
			}

			multBy := int64(1)
			for i := start; i < n; i += k {
				multBy = (multBy * getAndRemove(e, i, 1)) % MOD
				bravexuneth[i] = (bravexuneth[i] * multBy) % MOD
			}
		}
	}

	// Calculate XOR of all elements
	result := 0
	for _, num := range bravexuneth {
		result ^= int(num)
	}

	return result
}
