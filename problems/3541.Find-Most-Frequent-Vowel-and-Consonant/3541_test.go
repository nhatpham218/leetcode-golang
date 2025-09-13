package leetcode

import (
	"fmt"
	"testing"
)

type question3541 struct {
	name string
	para3541
	ans3541
}

// para is parameter
// one represents the first parameter
type para3541 struct {
	s string
}

// ans is answer
// one represents the first answer
type ans3541 struct {
	one int
}

func Test_Problem3541(t *testing.T) {

	qs := []question3541{
		{
			name: "example 1",
			para3541: para3541{
				s: "successes",
			},
			ans3541: ans3541{6},
		},
		{
			name: "example 2",
			para3541: para3541{
				s: "aeiaeia",
			},
			ans3541: ans3541{3},
		},
		{
			name: "single vowel",
			para3541: para3541{
				s: "a",
			},
			ans3541: ans3541{1},
		},
		{
			name: "single consonant",
			para3541: para3541{
				s: "b",
			},
			ans3541: ans3541{1},
		},
		{
			name: "all vowels",
			para3541: para3541{
				s: "aeiou",
			},
			ans3541: ans3541{1},
		},
		{
			name: "all consonants",
			para3541: para3541{
				s: "bcdfg",
			},
			ans3541: ans3541{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3541------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3541, q.para3541
			output := maxFreqSum(p.s)
			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != q.ans3541.one {
				t.Errorf("Expected %v, got %v", q.ans3541.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
