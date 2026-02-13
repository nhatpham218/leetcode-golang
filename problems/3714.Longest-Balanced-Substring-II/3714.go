package leetcode

// 3714. Longest Balanced Substring II
// https://leetcode.com/problems/longest-balanced-substring-ii/description/
func longestBalanced(s string) int {
	best := 0

	// Case 1: single character — longest run of the same character
	best = max(best, maxRun(s, 'a'))
	best = max(best, maxRun(s, 'b'))
	best = max(best, maxRun(s, 'c'))

	// Case 2: exactly two distinct characters — prefix difference (count_a - count_b) per pair
	best = max(best, twoChars(s, 'a', 'b'))
	best = max(best, twoChars(s, 'a', 'c'))
	best = max(best, twoChars(s, 'b', 'c'))

	// Case 3: all three characters — hash (count_b - count_a, count_c - count_a), same pair => balanced
	best = max(best, threeChars(s))

	return best
}

// maxRun returns the length of the longest run of character c in s.
func maxRun(s string, c byte) int {
	maxLen, cur := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cur++
			if cur > maxLen {
				maxLen = cur
			}
		} else {
			cur = 0
		}
	}
	return maxLen
}

// twoChars returns the longest balanced substring for the two given characters (ignore the third).
// Prefix diff d = count_a - count_b; same d at prefix end i and j => substring s[i+1..j] is balanced. Store earliest prefix-end index per d.
// When we see the third character, reset first so no substring includes it.
func twoChars(s string, a, b byte) int {
	type state struct{ idx, pa int }
	first := make(map[int]state)
	pa, pb := 0, 0
	first[0] = state{idx: -1, pa: 0} // before any char
	best := 0
	for j := 0; j < len(s); j++ {
		if s[j] == a {
			pa++
		} else if s[j] == b {
			pb++
		} else {
			// third character: substrings cannot include it; reset so next occurrence of d starts after j
			d := pa - pb
			first = map[int]state{d: {idx: j, pa: pa}}
			continue
		}
		d := pa - pb
		if st, ok := first[d]; ok {
			countA := pa - st.pa
			if countA >= 1 {
				if j-st.idx > best {
					best = j - st.idx
				}
			}
		} else {
			first[d] = state{idx: j, pa: pa}
		}
	}
	return best
}

// threeChars returns the longest balanced substring containing all three characters.
// Hash (count_b - count_a, count_c - count_a); same pair at prefix end i and j => substring s[i+1..j] has equal counts.
func threeChars(s string) int {
	type key struct{ ba, ca int }
	type state3 struct{ idx, pa int }
	first := make(map[key]state3)
	pa, pb, pc := 0, 0, 0
	first[key{0, 0}] = state3{idx: -1, pa: 0}
	best := 0
	for j := 0; j < len(s); j++ {
		switch s[j] {
		case 'a':
			pa++
		case 'b':
			pb++
		case 'c':
			pc++
		}
		k := key{pb - pa, pc - pa}
		if st, ok := first[k]; ok {
			countA := pa - st.pa
			if countA >= 1 {
				if j-st.idx > best {
					best = j - st.idx
				}
			}
		} else {
			first[k] = state3{idx: j, pa: pa}
		}
	}
	return best
}
