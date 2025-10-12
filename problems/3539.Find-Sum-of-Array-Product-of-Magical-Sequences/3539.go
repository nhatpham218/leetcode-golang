package leetcode

const MOD = 1000000007

// 3539. Find Sum of Array Product of Magical Sequences
// https://leetcode.com/problems/find-sum-of-array-product-of-magical-sequences/
func magicalSum(m int, k int, nums []int) int {
	n := len(nums)

	// Precompute factorials
	fac := make([]int, m+1)
	fac[0] = 1
	for i := 1; i <= m; i++ {
		fac[i] = (fac[i-1] * i) % MOD
	}

	// Precompute inverse factorials
	ifac := make([]int, m+1)
	ifac[0] = 1
	ifac[1] = 1
	for i := 2; i <= m; i++ {
		ifac[i] = modPow(i, MOD-2, MOD)
	}
	for i := 2; i <= m; i++ {
		ifac[i] = (ifac[i-1] * ifac[i]) % MOD
	}

	// Precompute powers of nums
	numsPower := make([][]int, n)
	for i := 0; i < n; i++ {
		numsPower[i] = make([]int, m+1)
		numsPower[i][0] = 1
		for j := 1; j <= m; j++ {
			numsPower[i][j] = (numsPower[i][j-1] * nums[i]) % MOD
		}
	}

	// DP: f[i][j][p][q]
	// i: current index (0 to n-1)
	// j: number of elements used (0 to m)
	// p: partial sum/carry (0 to m*2)
	// q: number of set bits counted so far (0 to k)
	maxP := m * 2
	f := make([][][][]int, n)
	for i := range f {
		f[i] = make([][][]int, m+1)
		for j := range f[i] {
			f[i][j] = make([][]int, maxP+1)
			for p := range f[i][j] {
				f[i][j][p] = make([]int, k+1)
			}
		}
	}

	// Initialize first position
	for j := 0; j <= m; j++ {
		f[0][j][j][0] = (numsPower[0][j] * ifac[j]) % MOD
	}

	// DP transitions
	for i := 0; i+1 < n; i++ {
		for j := 0; j <= m; j++ {
			for p := 0; p <= maxP; p++ {
				for q := 0; q <= k; q++ {
					if f[i][j][p][q] == 0 {
						continue
					}

					q2 := (p % 2) + q
					if q2 > k {
						break
					}

					for r := 0; r+j <= m; r++ {
						p2 := p/2 + r
						if p2 > maxP {
							break
						}
						contribution := f[i][j][p][q]
						contribution = (contribution * numsPower[i+1][r]) % MOD
						contribution = (contribution * ifac[r]) % MOD
						f[i+1][j+r][p2][q2] = (f[i+1][j+r][p2][q2] + contribution) % MOD
					}
				}
			}
		}
	}

	// Collect results
	result := 0
	for p := 0; p <= maxP; p++ {
		for q := 0; q <= k; q++ {
			if popCount(p)+q == k {
				contribution := (f[n-1][m][p][q] * fac[m]) % MOD
				result = (result + contribution) % MOD
			}
		}
	}

	return result
}

// popCount counts the number of set bits
func popCount(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// modInv calculates modular inverse using Fermat's little theorem
// a^(p-1) ≡ 1 (mod p) => a^(-1) ≡ a^(p-2) (mod p)
func modInv(a, mod int) int {
	return modPow(a, mod-2, mod)
}

// modPow calculates (base^exp) % mod using binary exponentiation
func modPow(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}
