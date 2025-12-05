package leetcode

import (
	"fmt"
	"testing"
)

type question3432 struct {
	name string
	para3432
	ans3432
}

// para is parameter
// one represents the first parameter
type para3432 struct {
	one []int
}

// ans is answer
// one represents the first answer
type ans3432 struct {
	one int
}

func Test_Problem3432(t *testing.T) {

	qs := []question3432{
		{
			name: "example 1",
			para3432: para3432{
				one: []int{10, 10, 3, 7, 6},
			},
			ans3432: ans3432{4},
		},
		{
			name: "example 2",
			para3432: para3432{
				one: []int{1, 2, 2},
			},
			ans3432: ans3432{0},
		},
		{
			name: "example 3",
			para3432: para3432{
				one: []int{2, 4, 6, 8},
			},
			ans3432: ans3432{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3432------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3432, q.para3432
			output := countPartitions(p.one)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.one, output)
			if output != q.ans3432.one {
				t.Errorf("Expected %v, got %v", q.ans3432.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

