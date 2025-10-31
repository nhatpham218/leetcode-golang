package leetcode

import (
	"fmt"
	"testing"
)

type question3289 struct {
	name string
	para3289
	ans3289
}

type para3289 struct {
	nums []int
}

type ans3289 struct {
	output []int
}

func Test_Problem3289(t *testing.T) {

	qs := []question3289{
		{
			name:     "Example 1",
			para3289: para3289{nums: []int{0, 1, 1, 0}},
			ans3289:  ans3289{output: []int{0, 1}},
		},
		{
			name:     "Example 2",
			para3289: para3289{nums: []int{0, 3, 2, 1, 3, 2}},
			ans3289:  ans3289{output: []int{2, 3}},
		},
		{
			name:     "Example 3",
			para3289: para3289{nums: []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}},
			ans3289:  ans3289{output: []int{4, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3289------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3289, q.para3289
			output := getSneakyNumbers(p.nums)

			fmt.Printf("[input]: %v       [output]:%v\n", p.nums, output)
			if !equalSlices(output, a.output) {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if !m[v] {
			return false
		}
	}
	return true
}
