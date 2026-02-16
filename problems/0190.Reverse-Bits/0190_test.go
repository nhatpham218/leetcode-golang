package leetcode

import (
	"fmt"
	"testing"
)

type question190 struct {
	name  string
	para190
	ans190
}

type para190 struct {
	one uint32
}

type ans190 struct {
	one uint32
}

func Test_Problem190(t *testing.T) {

	qs := []question190{
		{
			name:    "example 1",
			para190: para190{one: 43261596},
			ans190:  ans190{one: 964176192},
		},
		{
			name:    "example 2",
			para190: para190{one: 2147483644},
			ans190:  ans190{one: 1073741822},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 190------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans190, q.para190
			output := reverseBits(p.one)
			fmt.Printf("[input]: n=%v       [output]:%v\n", p.one, output)
			if output != q.ans190.one {
				t.Errorf("Expected %v, got %v", q.ans190.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
