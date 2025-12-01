package leetcode

import "sort"

// 3661. Maximum Walls Destroyed by Robots
// https://leetcode.com/problems/maximum-walls-destroyed-by-robots/description/
// There is an endless straight line populated with some robots and walls. You are given integer arrays robots, distance, and walls:
// robots[i] is the position of the ith robot.
// distance[i] is the maximum distance the ith robot's bullet can travel.
// walls[j] is the position of the jth wall.
// Every robot has one bullet that can either fire to the left or the right at most distance[i] meters.
//
// A bullet destroys every wall in its path that lies within its range. Robots are fixed obstacles: if a bullet hits another robot before reaching a wall, it immediately stops at that robot and cannot continue.
//
// Return the maximum number of unique walls that can be destroyed by the robots.
//
// Notes:
// - A wall and a robot may share the same position; the wall can be destroyed by the robot at that position.
// - Robots are not destroyed by bullets.
func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	if n == 0 {
		return 0
	}

	// Create array of [position, distance] pairs and sort by position
	type Robot struct {
		pos  int
		dist int
	}

	x := make([]Robot, n+1)
	for i := 0; i < n; i++ {
		x[i] = Robot{robots[i], distance[i]}
	}

	// Sort robots by position
	sort.Slice(x[:n], func(i, j int) bool {
		return x[i].pos < x[j].pos
	})

	// Add sentinel robot at the end
	x[n] = Robot{1e9, 0}

	// Sort walls
	sort.Ints(walls)

	// Helper function to count walls in range [l, r]
	query := func(l, r int) int {
		if l > r {
			return 0
		}
		// Find first wall >= l
		left := sort.SearchInts(walls, l)
		// Find first wall > r
		right := sort.SearchInts(walls, r+1)
		return right - left
	}

	// dp[i][0] = max walls destroyed when robot i shoots left
	// dp[i][1] = max walls destroyed when robot i shoots right
	dp := make([][2]int, n)

	// Base case for first robot
	dp[0][0] = query(x[0].pos-x[0].dist, x[0].pos)
	if n > 1 {
		rightLimit := min(x[1].pos-1, x[0].pos+x[0].dist)
		dp[0][1] = query(x[0].pos, rightLimit)
	} else {
		dp[0][1] = query(x[0].pos, x[0].pos+x[0].dist)
	}

	// Transition for remaining robots
	for i := 1; i < n; i++ {
		// Robot i shoots right
		rightLimit := min(x[i+1].pos-1, x[i].pos+x[i].dist)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]) + query(x[i].pos, rightLimit)

		// Robot i shoots left
		leftLimit := max(x[i].pos-x[i].dist, x[i-1].pos+1)
		dp[i][0] = dp[i-1][0] + query(leftLimit, x[i].pos)

		// Alternative case: previous robot shot right, current shoots left
		// Need to subtract overlap to avoid double counting
		prevRightLimit := min(x[i-1].pos+x[i-1].dist, x[i].pos-1)
		overlapLeft := max(leftLimit, x[i-1].pos+1)
		overlapRight := min(prevRightLimit, x[i].pos)

		res := dp[i-1][1] + query(leftLimit, x[i].pos) - query(overlapLeft, overlapRight)
		dp[i][0] = max(res, dp[i][0])
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
