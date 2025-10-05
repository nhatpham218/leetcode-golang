package leetcode

import (
	"fmt"
	"testing"
)

type question417 struct {
	name string
	para417
	ans417
}

// para is parameter
type para417 struct {
	heights [][]int
}

// ans is answer
type ans417 struct {
	one [][]int
}

func Test_Problem417(t *testing.T) {

	qs := []question417{
		{
			name: "example 1",
			para417: para417{
				heights: [][]int{
					{1, 2, 2, 3, 5},
					{3, 2, 3, 4, 4},
					{2, 4, 5, 3, 1},
					{6, 7, 1, 4, 5},
					{5, 1, 1, 2, 4},
				},
			},
			ans417: ans417{
				[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
			},
		},
		{
			name: "example 2",
			para417: para417{
				heights: [][]int{
					{1},
				},
			},
			ans417: ans417{
				[][]int{{0, 0}},
			},
		},
		{
			name: "single row",
			para417: para417{
				heights: [][]int{
					{1, 2, 3, 4, 5},
				},
			},
			ans417: ans417{
				[][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}},
			},
		},
		{
			name: "single column",
			para417: para417{
				heights: [][]int{
					{1},
					{2},
					{3},
					{4},
					{5},
				},
			},
			ans417: ans417{
				[][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
			},
		},
		{
			name: "simple 2x2",
			para417: para417{
				heights: [][]int{
					{1, 2},
					{2, 3},
				},
			},
			ans417: ans417{
				[][]int{{0, 1}, {1, 0}, {1, 1}},
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 417------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans417, q.para417

			output := pacificAtlantic(p.heights)
			fmt.Printf("[input]: heights=%v       [output]:%v\n", p.heights, output)
			if !isEqualResult(output, q.ans417.one) {
				t.Errorf("Expected %v, got %v", q.ans417.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

// Helper function to compare two result slices (order doesn't matter)
func isEqualResult(result1, result2 [][]int) bool {
	if len(result1) != len(result2) {
		return false
	}

	// Create maps to count occurrences
	count1 := make(map[string]int)
	count2 := make(map[string]int)

	for _, cell := range result1 {
		key := fmt.Sprintf("%d,%d", cell[0], cell[1])
		count1[key]++
	}

	for _, cell := range result2 {
		key := fmt.Sprintf("%d,%d", cell[0], cell[1])
		count2[key]++
	}

	// Check if counts match
	for key, count := range count1 {
		if count2[key] != count {
			return false
		}
	}

	for key, count := range count2 {
		if count1[key] != count {
			return false
		}
	}

	return true
}
