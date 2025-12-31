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
		{
			name: "example 5",
			para1970: para1970{
				row:   33,
				col:   3,
				cells: [][]int{{26, 1}, {1, 2}, {30, 2}, {10, 1}, {20, 1}, {23, 3}, {9, 1}, {30, 1}, {16, 1}, {2, 2}, {23, 2}, {31, 3}, {21, 1}, {21, 3}, {15, 3}, {28, 2}, {24, 2}, {5, 1}, {33, 1}, {18, 3}, {9, 2}, {16, 2}, {21, 2}, {14, 3}, {19, 2}, {1, 1}, {20, 2}, {2, 1}, {12, 3}, {2, 3}, {25, 2}, {26, 3}, {25, 3}, {13, 2}, {19, 3}, {29, 1}, {4, 2}, {27, 1}, {3, 2}, {17, 2}, {6, 3}, {17, 3}, {31, 1}, {27, 3}, {8, 2}, {24, 3}, {29, 2}, {16, 3}, {12, 1}, {9, 3}, {6, 2}, {10, 3}, {33, 2}, {22, 3}, {22, 2}, {7, 1}, {18, 1}, {32, 1}, {14, 1}, {32, 2}, {1, 3}, {18, 2}, {11, 3}, {12, 2}, {28, 1}, {19, 1}, {24, 1}, {30, 3}, {11, 2}, {4, 1}, {4, 3}, {20, 3}, {8, 1}, {23, 1}, {7, 3}, {27, 2}, {22, 1}, {26, 2}, {15, 2}, {14, 2}, {28, 3}, {13, 1}, {5, 2}, {10, 2}, {6, 1}, {33, 3}, {15, 1}, {13, 3}, {3, 3}, {3, 1}, {31, 2}, {11, 1}, {5, 3}, {8, 3}, {32, 3}, {17, 1}, {7, 2}, {29, 3}, {25, 1}},
			},
			ans1970: ans1970{0},
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
