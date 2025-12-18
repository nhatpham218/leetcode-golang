package leetcode

import (
	"fmt"
	"testing"
)

type question1913 struct {
	name string
	para1913
	ans1913
}

// para is parameter
// one represents the first parameter
type para1913 struct {
	one []int
}

// ans is answer
// one represents the first answer
type ans1913 struct {
	one int
}

func Test_Problem1913(t *testing.T) {

	qs := []question1913{
		{
			name: "example 1",
			para1913: para1913{
				one: []int{5, 6, 2, 7, 4},
			},
			ans1913: ans1913{34},
		},
		{
			name: "example 2",
			para1913: para1913{
				one: []int{4, 2, 5, 9, 7, 4, 8},
			},
			ans1913: ans1913{64},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1913------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1913, q.para1913
			output := maxProductDifference(p.one)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.one, output)
			if output != q.ans1913.one {
				t.Errorf("Expected %v, got %v", q.ans1913.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
