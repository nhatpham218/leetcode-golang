package leetcode

// 3021. Alice and Bob Playing Flower Game
// Alice and Bob are playing a turn-based game on a field, with two lanes of flowers between them.
// There are x flowers in the first lane and y flowers in the second lane.
// Alice takes the first turn. In each turn, a player picks one flower from either lane.
// The player who makes the last move wins.
// Return the number of pairs (x, y) where Alice wins, with x in [1,n] and y in [1,m].
func flowerGame(n int, m int) int64 {
	oddX := (n + 1) / 2
	evenX := n / 2

	oddY := (m + 1) / 2
	evenY := m / 2

	// Alice wins if total flowers (x + y) is odd
	// Alice wins if x is odd and y is even or x is even and y is odd
	// So we need to count the number of pairs where x is odd and y is even or x is even and y is odd
	return int64(oddX)*int64(evenY) + int64(evenX)*int64(oddY)
}
