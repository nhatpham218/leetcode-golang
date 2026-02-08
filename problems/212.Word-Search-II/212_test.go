package leetcode

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type question212 struct {
	name string
	para212
	ans212
}

// para is parameter
type para212 struct {
	board [][]byte
	words []string
}

// ans is answer
type ans212 struct {
	one []string
}

func Test_Problem212(t *testing.T) {
	qs := []question212{
		{
			name: "example 1",
			para212: para212{
				board: [][]byte{
					{'o', 'a', 'a', 'n'},
					{'e', 't', 'a', 'e'},
					{'i', 'h', 'k', 'r'},
					{'i', 'f', 'l', 'v'},
				},
				words: []string{"oath", "pea", "eat", "rain"},
			},
			ans212: ans212{[]string{"eat", "oath"}},
		},
		{
			name: "example 2",
			para212: para212{
				board: [][]byte{
					{'a', 'b'},
					{'c', 'd'},
				},
				words: []string{"abcb"},
			},
			ans212: ans212{[]string{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 212------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans212, q.para212
			output := findWords(p.board, p.words)
			if output == nil {
				output = []string{}
			}
			// Sort both for comparison since order doesn't matter
			sort.Strings(output)
			expected := make([]string, len(q.ans212.one))
			copy(expected, q.ans212.one)
			sort.Strings(expected)
			fmt.Printf("[input]: board=%v, words=%v       [output]:%v\n", p.board, p.words, output)
			if !reflect.DeepEqual(output, expected) {
				t.Errorf("Expected %v, got %v", q.ans212.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
