package leetcode

import (
	"fmt"
	"testing"
)

type question778 struct {
	name string
	para778
	ans778
}

// para is parameter
type para778 struct {
	grid [][]int
}

// ans is answer
type ans778 struct {
	one int
}

func Test_Problem778(t *testing.T) {

	qs := []question778{
		{
			name: "example 1",
			para778: para778{
				grid: [][]int{
					{0, 2},
					{1, 3},
				},
			},
			ans778: ans778{
				3,
			},
		},
		{
			name: "example 2",
			para778: para778{
				grid: [][]int{
					{0, 1, 2, 3, 4},
					{24, 23, 22, 21, 5},
					{12, 13, 14, 15, 16},
					{11, 17, 18, 19, 20},
					{10, 9, 8, 7, 6},
				},
			},
			ans778: ans778{
				16,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 778------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans778, q.para778

			output := swimInWater(p.grid)
			fmt.Printf("[input]: grid=%v       [output]:%v\n", p.grid, output)
			if output != q.ans778.one {
				t.Errorf("Expected %v, got %v", q.ans778.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
