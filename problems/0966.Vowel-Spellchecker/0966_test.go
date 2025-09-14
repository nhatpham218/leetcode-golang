package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question0966 struct {
	name string
	para0966
	ans0966
}

// para is parameter
// one represents the first parameter
type para0966 struct {
	wordlist []string
	queries  []string
}

// ans is answer
// one represents the first answer
type ans0966 struct {
	one []string
}

func Test_Problem0966(t *testing.T) {

	qs := []question0966{
		{
			name: "example 1",
			para0966: para0966{
				wordlist: []string{"KiTe", "kite", "hare", "Hare"},
				queries:  []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"},
			},
			ans0966: ans0966{[]string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"}},
		},
		{
			name: "example 2",
			para0966: para0966{
				wordlist: []string{"yellow"},
				queries:  []string{"YellOw"},
			},
			ans0966: ans0966{[]string{"yellow"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0966------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0966, q.para0966
			output := spellchecker(p.wordlist, p.queries)
			fmt.Printf("[input]: wordlist=%v, queries=%v       [output]:%v\n", p.wordlist, p.queries, output)
			if !reflect.DeepEqual(output, q.ans0966.one) {
				t.Errorf("Expected %v, got %v", q.ans0966.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
