package leetcode

import (
	"fmt"
	"testing"
)

type question2423 struct {
	para2423
	ans2423
}

// para is parameter
// one represents the first parameter
type para2423 struct {
	word string
}

// ans is answer
// one represents the first answer
type ans2423 struct {
	one bool
}

func Test_Problem2423(t *testing.T) {

	qs := []question2423{
		{
			para2423{"abcc"},
			ans2423{true},
		},
		{
			para2423{"aazz"},
			ans2423{false},
		},
		{
			para2423{"abc"},
			ans2423{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2423------------------------\n")

	for i, q := range qs {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			_, p := q.ans2423, q.para2423
			output := equalFrequency(p.word)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			if output != q.ans2423.one {
				t.Errorf("Expected %v, got %v", q.ans2423.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
