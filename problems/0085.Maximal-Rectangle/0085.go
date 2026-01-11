package leetcode

// 85. Maximal Rectangle
// https://leetcode.com/problems/maximal-rectangle/description/
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	cols := len(matrix[0])
	heights := make([]int, cols)
	result := 0

	for _, row := range matrix {
		for col := 0; col < cols; col++ {
			if row[col] == '1' {
				heights[col]++
			} else {
				heights[col] = 0
			}
		}
		result = max(result, largestRectangleInHistogram(heights))
	}

	return result
}

// largestRectangleInHistogram uses monotonic stack (LeetCode 84)
func largestRectangleInHistogram(heights []int) int {
	heights = append(heights, 0)
	stack := []int{}
	maxArea := 0

	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[top]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}

	return maxArea
}
