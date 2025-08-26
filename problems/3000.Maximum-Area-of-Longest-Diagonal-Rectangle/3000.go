package leetcode

// 3000. Maximum Area of Longest Diagonal Rectangle
func areaOfMaxDiagonal(dimensions [][]int) int {
	longestDiagonal := 0
	result := 0
	for _, dimension := range dimensions {
		diagonal := (dimension[0]*dimension[0] + dimension[1]*dimension[1])
		if diagonal > longestDiagonal {
			longestDiagonal = diagonal
			result = dimension[0] * dimension[1]
		} else if diagonal == longestDiagonal {
			result = max(result, dimension[0]*dimension[1])
		}
	}
	return result
}
