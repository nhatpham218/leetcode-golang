package leetcode

import (
	"fmt"
	"testing"
)

type question717 struct {
	name string
	para717
	ans717
}

type para717 struct {
	bits []int
}

type ans717 struct {
	output bool
}

func Test_Problem717(t *testing.T) {

	qs := []question717{
		{
			name: "Example 1",
			para717: para717{
				bits: []int{1, 0, 0},
			},
			ans717: ans717{output: true},
		},
		{
			name: "Example 2",
			para717: para717{
				bits: []int{1, 1, 1, 0},
			},
			ans717: ans717{output: false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 717------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans717, q.para717
			output := isOneBitCharacter(p.bits)

			fmt.Printf("[input]: bits=%v       [output]:%v\n", p.bits, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}


