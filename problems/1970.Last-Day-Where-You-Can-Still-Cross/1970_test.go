package leetcode

import (
	"fmt"
	"testing"
)

type question1970 struct {
	name string
	para1970
	ans1970
}

type para1970 struct {
	row   int
	col   int
	cells [][]int
}

type ans1970 struct {
	one int
}

func Test_Problem1970(t *testing.T) {

	qs := []question1970{
		{
			name: "example 1",
			para1970: para1970{
				row:   2,
				col:   2,
				cells: [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}},
			},
			ans1970: ans1970{2},
		},
		{
			name: "example 2",
			para1970: para1970{
				row:   2,
				col:   2,
				cells: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			},
			ans1970: ans1970{1},
		},
		{
			name: "example 3",
			para1970: para1970{
				row:   3,
				col:   3,
				cells: [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}},
			},
			ans1970: ans1970{3},
		},
		{
			name: "example 4",
			para1970: para1970{
				row:   5,
				col:   2,
				cells: [][]int{{5, 1}, {1, 2}, {3, 1}, {2, 2}, {3, 2}, {1, 1}, {5, 2}, {2, 1}, {4, 2}, {4, 1}},
			},
			ans1970: ans1970{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1970------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1970, q.para1970
			output := latestDayToCross(p.row, p.col, p.cells)
			fmt.Printf("[input]: row=%v, col=%v, cells=%v       [output]:%v\n", p.row, p.col, p.cells, output)
			if output != q.ans1970.one {
				t.Errorf("Expected %v, got %v", q.ans1970.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
