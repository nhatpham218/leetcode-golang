package leetcode

import (
	"fmt"
	"testing"
)

type question2197 struct {
	name string
	para2197
	ans2197
}

// para is parameter
// one represents the first parameter
type para2197 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans2197 struct {
	one []int
}

func Test_Problem2197(t *testing.T) {

	qs := []question2197{
		{
			name: "example 1",
			para2197: para2197{
				nums: []int{6, 4, 3, 2, 7, 6, 2},
			},
			ans2197: ans2197{[]int{12, 7, 6}},
		},
		{
			name: "example 2",
			para2197: para2197{
				nums: []int{2, 2, 1, 1, 3, 3, 3},
			},
			ans2197: ans2197{[]int{2, 1, 1, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2197------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2197, q.para2197
			output := replaceNonCoprimes(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if !sliceEqual(output, q.ans2197.one) {
				t.Errorf("Expected %v, got %v", q.ans2197.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

