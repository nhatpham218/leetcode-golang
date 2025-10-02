package leetcode

import (
	"fmt"
	"testing"
)

type question1518 struct {
	name string
	para1518
	ans1518
}

// para is parameter
// one represents the first parameter
// two represents the second parameter
type para1518 struct {
	numBottles  int
	numExchange int
}

// ans is answer
// one represents the first answer
type ans1518 struct {
	one int
}

func Test_Problem1518(t *testing.T) {

	qs := []question1518{
		{
			name: "example 1",
			para1518: para1518{
				numBottles:  9,
				numExchange: 3,
			},
			ans1518: ans1518{13},
		},
		{
			name: "example 2",
			para1518: para1518{
				numBottles:  15,
				numExchange: 4,
			},
			ans1518: ans1518{19},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1518------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1518, q.para1518
			output := numWaterBottles(p.numBottles, p.numExchange)
			fmt.Printf("[input]: numBottles=%d, numExchange=%d       [output]:%d\n", p.numBottles, p.numExchange, output)
			if output != q.ans1518.one {
				t.Errorf("Expected %d, got %d", q.ans1518.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
