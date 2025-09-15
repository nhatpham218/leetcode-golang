package leetcode

import (
	"fmt"
	"testing"
)

type question1935 struct {
	name string
	para1935
	ans1935
}

// para is parameter
// one represents the first parameter
type para1935 struct {
	text          string
	brokenLetters string
}

// ans is answer
// one represents the first answer
type ans1935 struct {
	one int
}

func Test_Problem1935(t *testing.T) {

	qs := []question1935{
		{
			name: "example 1",
			para1935: para1935{
				text:          "hello world",
				brokenLetters: "ad",
			},
			ans1935: ans1935{1},
		},
		{
			name: "example 2",
			para1935: para1935{
				text:          "leet code",
				brokenLetters: "lt",
			},
			ans1935: ans1935{1},
		},
		{
			name: "example 3",
			para1935: para1935{
				text:          "leet code",
				brokenLetters: "e",
			},
			ans1935: ans1935{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1935------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1935, q.para1935
			output := canBeTypedWords(p.text, p.brokenLetters)
			fmt.Printf("[input]: text=%v, brokenLetters=%v       [output]:%v\n", p.text, p.brokenLetters, output)
			if output != q.ans1935.one {
				t.Errorf("Expected %v, got %v", q.ans1935.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
