package leetcode

import (
	"fmt"
	"testing"
)

type question744 struct {
	name string
	para744
	ans744
}

// para is parameter
// one represents the first parameter
type para744 struct {
	letters []byte
	target  byte
}

// ans is answer
// one represents the first answer
type ans744 struct {
	one byte
}

func Test_Problem744(t *testing.T) {

	qs := []question744{
		{
			name: "example 1",
			para744: para744{
				letters: []byte{'c', 'f', 'j'},
				target:  'a',
			},
			ans744: ans744{'c'},
		},
		{
			name: "example 2",
			para744: para744{
				letters: []byte{'c', 'f', 'j'},
				target:  'c',
			},
			ans744: ans744{'f'},
		},
		{
			name: "example 3",
			para744: para744{
				letters: []byte{'x', 'x', 'y', 'y'},
				target:  'z',
			},
			ans744: ans744{'x'},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 744------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans744, q.para744
			output := nextGreatestLetter(p.letters, p.target)
			fmt.Printf("[input]: letters=%v, target=%c       [output]:%c\n", p.letters, p.target, output)
			if output != q.ans744.one {
				t.Errorf("Expected %c, got %c", q.ans744.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
