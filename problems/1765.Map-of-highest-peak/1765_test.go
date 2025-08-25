package leetcode

import (
	"reflect"
	"testing"
)

func Test_highestPeak(t *testing.T) {
	type args struct {
		isWater [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1 - 2x2 grid with one water cell",
			args: args{
				isWater: [][]int{{0, 1}, {0, 0}},
			},
			want: [][]int{{1, 0}, {2, 1}},
		},
		{
			name: "Example 2 - 3x3 grid with two water cells",
			args: args{
				isWater: [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}},
			},
			want: [][]int{{1, 1, 0}, {0, 1, 1}, {1, 2, 2}},
		},
		{
			name: "Single cell - water",
			args: args{
				isWater: [][]int{{1}},
			},
			want: [][]int{{0}},
		},
		{
			name: "Single cell - land",
			args: args{
				isWater: [][]int{{0}},
			},
			want: [][]int{{1}},
		},
		{
			name: "All water cells",
			args: args{
				isWater: [][]int{{1, 1}, {1, 1}},
			},
			want: [][]int{{0, 0}, {0, 0}},
		},
		{
			name: "All land cells",
			args: args{
				isWater: [][]int{{0, 0}, {0, 0}},
			},
			want: [][]int{{1, 1}, {1, 1}},
		},
		{
			name: "Water in center",
			args: args{
				isWater: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			},
			want: [][]int{{2, 1, 2}, {1, 0, 1}, {2, 1, 2}},
		},
		{
			name: "Water on edge",
			args: args{
				isWater: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			},
			want: [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}},
		},
		{
			name: "Multiple water cells scattered",
			args: args{
				isWater: [][]int{{1, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 1}},
			},
			want: [][]int{{0, 1, 1, 0}, {1, 2, 1, 1}, {1, 2, 1, 1}, {0, 1, 1, 0}},
		},
		{
			name: "Linear arrangement - horizontal",
			args: args{
				isWater: [][]int{{1, 0, 0, 0, 1}},
			},
			want: [][]int{{0, 1, 2, 1, 0}},
		},
		{
			name: "Linear arrangement - vertical",
			args: args{
				isWater: [][]int{{1}, {0}, {0}, {0}, {1}},
			},
			want: [][]int{{0}, {1}, {2}, {1}, {0}},
		},
		{
			name: "Checkered pattern",
			args: args{
				isWater: [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}},
			},
			want: [][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := highestPeak(tt.args.isWater)
			if !isValidHeightMatrix(got, tt.args.isWater) {
				t.Errorf("highestPeak() returned invalid height matrix = %v", got)
				return
			}
			// For this problem, there can be multiple valid solutions
			// We check if the solution is valid rather than exact match
			if !reflect.DeepEqual(got, tt.want) {
				// Check if it's an alternative valid solution
				if !hasValidAlternativeSolution(got, tt.want, tt.args.isWater) {
					t.Errorf("highestPeak() = %v, want %v or equivalent valid solution", got, tt.want)
				}
			}
		})
	}
}

// Helper function to validate if the height matrix satisfies all constraints
func isValidHeightMatrix(height [][]int, isWater [][]int) bool {
	if len(height) != len(isWater) || len(height[0]) != len(isWater[0]) {
		return false
	}

	m, n := len(height), len(height[0])

	// Check all constraints
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Height must be non-negative
			if height[i][j] < 0 {
				return false
			}

			// Water cells must have height 0
			if isWater[i][j] == 1 && height[i][j] != 0 {
				return false
			}

			// Check adjacent cells constraint
			directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					if abs(height[i][j]-height[ni][nj]) > 1 {
						return false
					}
				}
			}
		}
	}

	return true
}

// Helper function to check if the solution has similar maximum height as expected
func hasValidAlternativeSolution(got, want [][]int, isWater [][]int) bool {
	if !isValidHeightMatrix(got, isWater) {
		return false
	}

	// Check if maximum heights are the same
	maxGot, maxWant := getMaxHeight(got), getMaxHeight(want)
	return maxGot == maxWant
}

func getMaxHeight(matrix [][]int) int {
	maxHeight := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] > maxHeight {
				maxHeight = matrix[i][j]
			}
		}
	}
	return maxHeight
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
