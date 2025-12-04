package leetcode

// 2211. Count Collisions on a Road
// https://leetcode.com/problems/count-collisions-on-a-road/description/
func countCollisions(directions string) int {
	n := len(directions)
	ans := 0
	accident := false

	for i := 1; i < n; i++ {
		if directions[i] == 'L' {
			if directions[i-1] != 'L' || accident {
				ans++
				accident = true
			}
		}
	}

	accident = false
	for i := n - 2; i >= 0; i-- {
		if directions[i] == 'R' {
			if directions[i+1] != 'R' || accident {
				ans++
				accident = true
			}
		}
	}

	return ans
}
