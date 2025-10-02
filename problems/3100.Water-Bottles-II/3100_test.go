package leetcode

import (
	"fmt"
	"testing"
)

type question3100 struct {
	name string
	para3100
	ans3100
}

// para is parameter
// one represents the first parameter
// two represents the second parameter
type para3100 struct {
	numBottles  int
	numExchange int
}

// ans is answer
// one represents the first answer
type ans3100 struct {
	one int
}

func Test_Problem3100(t *testing.T) {

	qs := []question3100{
		{
			name: "example 1",
			para3100: para3100{
				numBottles:  13,
				numExchange: 6,
			},
			ans3100: ans3100{15},
		},
		{
			name: "example 2",
			para3100: para3100{
				numBottles:  10,
				numExchange: 3,
			},
			ans3100: ans3100{13},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3100------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3100, q.para3100
			output := maxBottlesDrunk(p.numBottles, p.numExchange)
			fmt.Printf("[input]: numBottles=%d, numExchange=%d       [output]:%d\n", p.numBottles, p.numExchange, output)
			if output != q.ans3100.one {
				t.Errorf("Expected %d, got %d", q.ans3100.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
