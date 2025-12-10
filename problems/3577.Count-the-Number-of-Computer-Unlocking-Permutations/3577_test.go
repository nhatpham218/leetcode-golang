package leetcode

import (
	"fmt"
	"testing"
)

type question3577 struct {
	name string
	para3577
	ans3577
}

// para is parameter
// one represents the first parameter
type para3577 struct {
	complexity []int
}

// ans is answer
// one represents the first answer
type ans3577 struct {
	one int
}

func Test_Problem3577(t *testing.T) {

	qs := []question3577{
		{
			name: "example 1",
			para3577: para3577{
				complexity: []int{1, 2, 3},
			},
			ans3577: ans3577{2},
		},
		{
			name: "example 2",
			para3577: para3577{
				complexity: []int{3, 3, 3, 4, 4, 4},
			},
			ans3577: ans3577{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3577------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3577, q.para3577
			output := countPermutations(p.complexity)
			fmt.Printf("[input]: complexity=%v       [output]:%v\n", p.complexity, output)
			if output != q.ans3577.one {
				t.Errorf("Expected %v, got %v", q.ans3577.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
