package leetcode

import (
	"fmt"
	"testing"
)

type question3354 struct {
	name string
	para3354
	ans3354
}

type para3354 struct {
	nums []int
}

type ans3354 struct {
	output int
}

func Test_Problem3354(t *testing.T) {

	qs := []question3354{}

	fmt.Printf("------------------------Leetcode Problem 3354------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3354, q.para3354
			output := countValidSelections(p.nums)

			fmt.Printf("[input]: %v       [output]:%v\n", p.nums, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
