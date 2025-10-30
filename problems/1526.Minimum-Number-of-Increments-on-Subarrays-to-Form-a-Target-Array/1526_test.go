package leetcode

import (
	"fmt"
	"testing"
)

type question1526 struct {
	name string
	para1526
	ans1526
}

type para1526 struct {
	target []int
}

type ans1526 struct {
	output int
}

func Test_Problem1526(t *testing.T) {

	qs := []question1526{
		{
			name:     "Example 1",
			para1526: para1526{target: []int{1, 2, 3, 2, 1}},
			ans1526:  ans1526{output: 3},
		},
		{
			name:     "Example 2",
			para1526: para1526{target: []int{3, 1, 1, 2}},
			ans1526:  ans1526{output: 4},
		},
		{
			name:     "Example 3",
			para1526: para1526{target: []int{3, 1, 5, 4, 2}},
			ans1526:  ans1526{output: 7},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1526------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans1526, q.para1526
			output := minNumberOperations(p.target)

			fmt.Printf("[input]: %v       [output]:%v\n", p.target, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
