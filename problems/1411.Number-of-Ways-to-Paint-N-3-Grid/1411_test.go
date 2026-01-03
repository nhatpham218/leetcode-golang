package leetcode

import (
	"fmt"
	"testing"
)

type question1411 struct {
	name string
	para1411
	ans1411
}

type para1411 struct {
	n int
}

type ans1411 struct {
	one int
}

func Test_Problem1411(t *testing.T) {

	qs := []question1411{
		{
			name:     "example 1",
			para1411: para1411{n: 1},
			ans1411:  ans1411{12},
		},
		{
			name:     "example 2",
			para1411: para1411{n: 5000},
			ans1411:  ans1411{30228214},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1411------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1411, q.para1411
			output := numOfWays(p.n)
			fmt.Printf("[input]: n=%v       [output]:%v\n", p.n, output)
			if output != q.ans1411.one {
				t.Errorf("Expected %v, got %v", q.ans1411.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

