package leetcode

// 3721. Longest Balanced Subarray II
// https://leetcode.com/problems/longest-balanced-subarray-ii/description/
func longestBalanced(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// sign: +1 for odd, -1 for even
	sign := func(v int) int {
		if v%2 != 0 {
			return 1
		}
		return -1
	}

	// nextOcc[i] = next index where nums[i] appears, or n
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}
	nextOcc := make([]int, n)
	for _, indices := range pos {
		for j := 0; j < len(indices); j++ {
			if j+1 < len(indices) {
				nextOcc[indices[j]] = indices[j+1]
			} else {
				nextOcc[indices[j]] = n
			}
		}
	}

	// firstOcc[v] = first index with value v
	firstOcc := make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		firstOcc[nums[i]] = i
	}

	// Lazy segment tree over [0..n-1]: range add, query rightmost 0 in [ql..qr]
	type seg struct {
		n     int
		minV  []int
		maxV  []int
		lazy  []int
	}

	newSeg := func(size int) *seg {
		s := &seg{n: size, minV: make([]int, 4*size), maxV: make([]int, 4*size), lazy: make([]int, 4*size)}
		return s
	}

	var update func(*seg, int, int, int, int, int, int)
	update = func(s *seg, idx, lo, hi, l, r, d int) {
		if r < lo || l > hi {
			return
		}
		if l <= lo && hi <= r {
			s.lazy[idx] += d
			s.minV[idx] += d
			s.maxV[idx] += d
			return
		}
		mid := (lo + hi) / 2
		// push
		if s.lazy[idx] != 0 {
			s.lazy[2*idx+1] += s.lazy[idx]
			s.minV[2*idx+1] += s.lazy[idx]
			s.maxV[2*idx+1] += s.lazy[idx]
			s.lazy[2*idx+2] += s.lazy[idx]
			s.minV[2*idx+2] += s.lazy[idx]
			s.maxV[2*idx+2] += s.lazy[idx]
			s.lazy[idx] = 0
		}
		update(s, 2*idx+1, lo, mid, l, r, d)
		update(s, 2*idx+2, mid+1, hi, l, r, d)
		s.minV[idx] = min(s.minV[2*idx+1], s.minV[2*idx+2])
		s.maxV[idx] = max(s.maxV[2*idx+1], s.maxV[2*idx+2])
	}

	var queryRightmostZero func(*seg, int, int, int, int, int) int
	queryRightmostZero = func(s *seg, idx, lo, hi, ql, qr int) int {
		if qr < lo || ql > hi {
			return -1
		}
		if s.minV[idx] > 0 || s.maxV[idx] < 0 {
			return -1
		}
		if lo == hi {
			if s.minV[idx] == 0 {
				return lo
			}
			return -1
		}
		mid := (lo + hi) / 2
		if s.lazy[idx] != 0 {
			s.lazy[2*idx+1] += s.lazy[idx]
			s.minV[2*idx+1] += s.lazy[idx]
			s.maxV[2*idx+1] += s.lazy[idx]
			s.lazy[2*idx+2] += s.lazy[idx]
			s.minV[2*idx+2] += s.lazy[idx]
			s.maxV[2*idx+2] += s.lazy[idx]
			s.lazy[idx] = 0
		}
		// right first to get rightmost
		if r := queryRightmostZero(s, 2*idx+2, mid+1, hi, ql, qr); r >= 0 {
			return r
		}
		return queryRightmostZero(s, 2*idx+1, lo, mid, ql, qr)
	}

	segTree := newSeg(n)

	// Initialize: for each value v, add sign(v) to [firstOcc[v], n-1]
	for v, p := range firstOcc {
		update(segTree, 0, 0, n-1, p, n-1, sign(v))
	}

	ans := 0
	for l := 0; l < n; l++ {
		// query for any r >= l with value 0
		if r := queryRightmostZero(segTree, 0, 0, n-1, l, n-1); r >= 0 {
			if r-l+1 > ans {
				ans = r - l + 1
			}
		}
		// pop pos[nums[l]]: remove contribution of nums[l] from [0, next-1]
		next := nextOcc[l]
		if next > 0 {
			update(segTree, 0, 0, n-1, 0, next-1, -sign(nums[l]))
		}
	}
	return ans
}
