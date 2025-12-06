package leetcode

import (
	"fmt"
	"testing"
)

type question3578 struct {
	name string
	para3578
	ans3578
}

// para is parameter
// one represents the first parameter
// two represents the second parameter

type para3578 struct {
	one []int
	two int
}

// ans is answer
// one represents the first answer

type ans3578 struct {
	one int
}

func Test_Problem3578(t *testing.T) {

	qs := []question3578{
		{
			name: "example 1",
			para3578: para3578{
				one: []int{9, 4, 1, 3, 7},
				two: 4,
			},
			ans3578: ans3578{6},
		},
		{
			name: "example 2",
			para3578: para3578{
				one: []int{3, 3, 4},
				two: 0,
			},
			ans3578: ans3578{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3578------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3578, q.para3578
			output := countPartitions(p.one, p.two)
			fmt.Printf("[input]: nums=%v, k=%v       [output]:%v\n", p.one, p.two, output)
			if output != q.ans3578.one {
				t.Errorf("Expected %v, got %v", q.ans3578.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
