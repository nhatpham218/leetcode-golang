package leetcode

// 3625. Count Number of Trapezoids II
// https://leetcode.com/problems/count-number-of-trapezoids-ii/description/
func countTrapezoids(points [][]int) int {
	slopeLines := make(map[int]map[int]int)
	vectorLines := make(map[int]map[int]int)
	n := len(points)

	addLine := func(mp map[int]map[int]int, key int, intercept int) {
		if _, ok := mp[key]; !ok {
			mp[key] = make(map[int]int)
		}
		mp[key][intercept]++
	}

	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dx := x2 - x1
			dy := y2 - y1

			if dx < 0 || (dx == 0 && dy < 0) {
				dx = -dx
				dy = -dy
			}

			g := gcd(dx, abs(dy))
			slopeX := dx / g
			slopeY := dy / g

			intercept := slopeX*y1 - slopeY*x1

			slopeKey := (slopeX << 12) | (slopeY + 2000)
			vectorKey := (dx << 12) | (dy + 2000)

			addLine(slopeLines, slopeKey, intercept)
			addLine(vectorLines, vectorKey, intercept)
		}
	}

	return countPairs(slopeLines) - countPairs(vectorLines)/2
}

func gcd(a, b int) int {
	a = abs(a)
	b = abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func countPairs(mp map[int]map[int]int) int {
	ans := 0
	for _, inner := range mp {
		total := 0
		for _, cnt := range inner {
			total += cnt
		}
		rem := total
		for _, cnt := range inner {
			rem -= cnt
			ans += cnt * rem
		}
	}
	return ans
}
