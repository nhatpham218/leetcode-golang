package leetcode

import (
	"fmt"
	"testing"
)

type question3539 struct {
	name string
	para3539
	ans3539
}

// para is parameter
// one represents the first parameter
type para3539 struct {
	m    int
	k    int
	nums []int
}

// ans is answer
// one represents the first answer
type ans3539 struct {
	one int
}

func Test_Problem3539(t *testing.T) {

	qs := []question3539{
		{
			name: "example 1",
			para3539: para3539{
				m:    5,
				k:    5,
				nums: []int{1, 10, 100, 10000, 1000000},
			},
			ans3539: ans3539{991600007},
		},
		{
			name: "example 2",
			para3539: para3539{
				m:    2,
				k:    2,
				nums: []int{5, 4, 3, 2, 1},
			},
			ans3539: ans3539{170},
		},
		{
			name: "example 3",
			para3539: para3539{
				m:    1,
				k:    1,
				nums: []int{28},
			},
			ans3539: ans3539{28},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3539------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3539, q.para3539
			output := magicalSum(p.m, p.k, p.nums)
			fmt.Printf("[input]: m=%v k=%v nums=%v       [output]:%v\n", p.m, p.k, p.nums, output)
			if output != q.ans3539.one {
				t.Errorf("Expected %v, got %v", q.ans3539.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
