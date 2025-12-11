package leetcode

import (
	"fmt"
	"testing"
)

type question3531 struct {
	name string
	para3531
	ans3531
}

// para is parameter
// one represents the first parameter
type para3531 struct {
	n         int
	buildings [][]int
}

// ans is answer
// one represents the first answer
type ans3531 struct {
	one int
}

func Test_Problem3531(t *testing.T) {

	qs := []question3531{
		{
			name: "example 1",
			para3531: para3531{
				n:         3,
				buildings: [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}},
			},
			ans3531: ans3531{1},
		},
		{
			name: "example 2",
			para3531: para3531{
				n:         3,
				buildings: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			},
			ans3531: ans3531{0},
		},
		{
			name: "example 3",
			para3531: para3531{
				n:         5,
				buildings: [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}},
			},
			ans3531: ans3531{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3531------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3531, q.para3531
			output := countCoveredBuildings(p.n, p.buildings)
			fmt.Printf("[input]: n=%v, buildings=%v       [output]:%v\n", p.n, p.buildings, output)
			if output != q.ans3531.one {
				t.Errorf("Expected %v, got %v", q.ans3531.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
