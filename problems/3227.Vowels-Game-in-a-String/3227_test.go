package leetcode

import (
	"fmt"
	"testing"
)

type question3227 struct {
	name string
	para3227
	ans3227
}

// para is parameter
// one represents the first parameter
type para3227 struct {
	s string
}

// ans is answer
// one represents the first answer
type ans3227 struct {
	one bool
}

func Test_Problem3227(t *testing.T) {

	qs := []question3227{
		{
			name: "example 1",
			para3227: para3227{
				s: "leetcoder",
			},
			ans3227: ans3227{true},
		},
		{
			name: "example 2",
			para3227: para3227{
				s: "bbcd",
			},
			ans3227: ans3227{false},
		},
		{
			name: "single vowel",
			para3227: para3227{
				s: "a",
			},
			ans3227: ans3227{true},
		},
		{
			name: "single consonant",
			para3227: para3227{
				s: "b",
			},
			ans3227: ans3227{false},
		},
		{
			name: "all vowels",
			para3227: para3227{
				s: "aeiou",
			},
			ans3227: ans3227{true},
		},
		{
			name: "all consonants",
			para3227: para3227{
				s: "bcdfg",
			},
			ans3227: ans3227{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3227------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3227, q.para3227
			output := doesAliceWin(p.s)
			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != q.ans3227.one {
				t.Errorf("Expected %v, got %v", q.ans3227.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
