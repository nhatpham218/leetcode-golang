package leetcode

import (
	"fmt"
	"testing"
)

type question407 struct {
	name string
	para407
	ans407
}

// para is parameter
// one represents the first parameter
type para407 struct {
	heightMap [][]int
}

// ans is answer
// one represents the first answer
type ans407 struct {
	one int
}

func Test_Problem407(t *testing.T) {

	qs := []question407{
		{
			name: "example 1",
			para407: para407{
				heightMap: [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}},
			},
			ans407: ans407{4},
		},
		{
			name: "example 2",
			para407: para407{
				heightMap: [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}},
			},
			ans407: ans407{10},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 407------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans407, q.para407
			output := trapRainWater(p.heightMap)
			fmt.Printf("[input]: heightMap=%v       [output]:%d\n", p.heightMap, output)
			if output != q.ans407.one {
				t.Errorf("Expected %d, got %d", q.ans407.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
