package leetcode

import (
	"fmt"
	"testing"
)

type question3397 struct {
	name string
	para3397
	ans3397
}

// para is parameter
// one represents the first parameter
type para3397 struct {
	nums []int
	k    int
}

// ans is answer
// one represents the first answer
type ans3397 struct {
	one int
}

func Test_Problem3397(t *testing.T) {

	qs := []question3397{
		{
			name: "example 1",
			para3397: para3397{
				nums: []int{1, 2, 2, 3, 3, 4},
				k:    2,
			},
			ans3397: ans3397{6},
		},
		{
			name: "example 2",
			para3397: para3397{
				nums: []int{4, 4, 4, 4},
				k:    1,
			},
			ans3397: ans3397{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3397------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3397, q.para3397
			output := maxDistinctElements(p.nums, p.k)
			fmt.Printf("[input]: nums=%v k=%v       [output]:%v\n", p.nums, p.k, output)
			if output != q.ans3397.one {
				t.Errorf("Expected %v, got %v", q.ans3397.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
