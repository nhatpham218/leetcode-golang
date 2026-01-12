package leetcode

// 1266. Minimum Time Visiting All Points
// https://leetcode.com/problems/minimum-time-visiting-all-points/description/
func minTimeToVisitAllPoints(points [][]int) int {
	time := 0
	for i := 1; i < len(points); i++ {
		time += max(abs(points[i][0]-points[i-1][0]), abs(points[i][1]-points[i-1][1]))
	}
	return time
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
