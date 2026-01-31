package leetcode

import (
	"fmt"
	"testing"
)

type question756 struct {
	name string
	para756
	ans756
}

type para756 struct {
	bottom  string
	allowed []string
}

type ans756 struct {
	one bool
}

func Test_Problem756(t *testing.T) {

	qs := []question756{
		{
			name: "example 1",
			para756: para756{
				bottom:  "BCD",
				allowed: []string{"BCC", "CDE", "CEA", "FFF"},
			},
			ans756: ans756{true},
		},
		{
			name: "example 2",
			para756: para756{
				bottom:  "AAAA",
				allowed: []string{"AAB", "AAC", "BCD", "BBE", "DEF"},
			},
			ans756: ans756{false},
		},
		{
			name: "example 3",
			para756: para756{
				bottom:  "ABCD",
				allowed: []string{"ABC", "BCA", "CDA", "ABD", "BCE", "CDF", "DEA", "EFF", "AFF"},
			},
			ans756: ans756{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 756------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans756, q.para756
			output := pyramidTransition(p.bottom, p.allowed)
			fmt.Printf("[input]: bottom=%v, allowed=%v       [output]:%v\n", p.bottom, p.allowed, output)
			if output != q.ans756.one {
				t.Errorf("Expected %v, got %v", q.ans756.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
