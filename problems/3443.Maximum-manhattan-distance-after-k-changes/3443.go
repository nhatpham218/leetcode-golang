package leetcode

import "math"

// maxDistance finds the maximum Manhattan distance from origin achievable at any point
// during the journey with the ability to change at most k characters in the string s
func maxDistance(s string, k int) int {
	ans := 0
	north, south, east, west := 0, 0, 0, 0

	// Helper function to calculate the maximum distance on one axis after modifications
	// drt1, drt2: counts of moves in opposite directions (e.g., north vs south)
	// times: number of changes we can make to optimize this axis
	count := func(drt1, drt2, times int) int {
		// Current net distance is |drt1 - drt2|
		// Each change can flip a move from one direction to another, gaining +2 distance
		// So we can improve by times*2 at most
		return int(math.Abs(float64(drt1-drt2))) + times*2
	} // Calculate modified Manhattan distance

	// Process each character in the string, tracking cumulative position
	for _, it := range s {
		// Update direction counters based on current move
		switch it {
		case 'N':
			north++
		case 'S':
			south++
		case 'E':
			east++
		case 'W':
			west++
		}

		// Greedily allocate changes to maximize distance:
		// First, use changes to eliminate opposing moves on N-S axis
		// The maximum useful changes for N-S is min(north, south) since we can only
		// flip moves that exist in the minority direction
		times1 := int(math.Min(float64(math.Min(float64(north), float64(south))), float64(k))) // modification times for N and S

		// Use remaining changes for E-W axis optimization
		times2 := int(math.Min(float64(math.Min(float64(east), float64(west))), float64(k)-float64(times1))) // modification times for E and W

		// Calculate total distance: optimized N-S distance + optimized E-W distance
		current := count(north, south, times1) + count(east, west, times2)

		// Track maximum distance seen so far
		if current > ans {
			ans = current
		}
	}

	return ans
}
