package leetcode

import (
	"fmt"
	"testing"
)

type question2125 struct {
	name string
	para2125
	ans2125
}

type para2125 struct {
	bank []string
}

type ans2125 struct {
	output int
}

func Test_Problem2125(t *testing.T) {

	qs := []question2125{}

	fmt.Printf("------------------------Leetcode Problem 2125------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans2125, q.para2125
			output := numberOfBeams(p.bank)

			fmt.Printf("[input]: %v       [output]:%v\n", p.bank, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
