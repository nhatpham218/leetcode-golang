package leetcode

import (
	"fmt"
	"testing"
)

type question0134 struct {
	name string
	para0134
	ans0134
}

// para is parameter
// one represents the first parameter
type para0134 struct {
	gas  []int
	cost []int
}

// ans is answer
// one represents the first answer
type ans0134 struct {
	one int
}

func Test_Problem0134(t *testing.T) {

	qs := []question0134{
		{
			name: "example 1",
			para0134: para0134{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			ans0134: ans0134{3},
		},
		{
			name: "example 2",
			para0134: para0134{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			ans0134: ans0134{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0134------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0134, q.para0134
			output := canCompleteCircuit(p.gas, p.cost)
			fmt.Printf("[input]: gas=%v, cost=%v       [output]:%v\n", p.gas, p.cost, output)
			if output != q.ans0134.one {
				t.Errorf("Expected %v, got %v", q.ans0134.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
