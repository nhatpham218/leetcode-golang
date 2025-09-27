package leetcode

import "math"

// 812. Largest Triangle Area
// https://leetcode.com/problems/largest-triangle-area
func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	ans := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				ans = max(ans, area(points[i], points[j], points[k]))
			}
		}
	}
	return ans
}

func area(p, q, r []int) float64 {
	return 0.5 * math.Abs(float64(p[0]*q[1]+q[0]*r[1]+r[0]*p[1]-p[1]*q[0]-q[1]*r[0]-r[1]*p[0]))
}
