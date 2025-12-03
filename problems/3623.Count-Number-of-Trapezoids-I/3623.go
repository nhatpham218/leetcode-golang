package leetcode

// 3623. Count Number of Trapezoids I
// https://leetcode.com/problems/count-number-of-trapezoids-i/description/
func countTrapezoids(points [][]int) int {
	const MOD = 1000000007
	// group points by y coordinate
	pointsByY := make(map[int]int)
	for _, point := range points {
		pointsByY[point[1]] += 1
	}

	res := 0
	edges := 0
	for _, points := range pointsByY {
		res = (res + edges*choose2(points)) % MOD
		edges += choose2(points)
	}

	return res
}

func choose2(n int) int {
	if n < 2 {
		return 0
	}
	return n * (n - 1) / 2
}
