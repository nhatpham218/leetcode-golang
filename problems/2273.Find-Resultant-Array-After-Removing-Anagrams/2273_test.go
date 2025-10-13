package leetcode

import (
	"fmt"
	"testing"
)

type question2273 struct {
	name string
	para2273
	ans2273
}

type para2273 struct {
	words []string
}

type ans2273 struct {
	one []string
}

func Test_Problem2273(t *testing.T) {

	qs := []question2273{}

	fmt.Printf("------------------------Leetcode Problem 2273------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2273, q.para2273
			output := removeAnagrams(p.words)
			fmt.Printf("[input]: words=%v       [output]:%v\n", p.words, output)
		})
	}
	fmt.Printf("\n\n\n")
}
