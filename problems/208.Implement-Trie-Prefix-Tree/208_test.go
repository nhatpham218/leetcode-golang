package leetcode

import (
	"fmt"
	"testing"
)

type question208 struct {
	name string
	para208
	ans208
}

type para208 struct {
	operations []string
	arguments  [][]string
}

type ans208 struct {
	expected []interface{}
}

func Test_Problem208(t *testing.T) {
	qs := []question208{
		{
			name: "example 1",
			para208: para208{
				operations: []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
				arguments:  [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}},
			},
			ans208: ans208{
				expected: []interface{}{nil, nil, true, false, true, nil, true},
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 208------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			trie := Constructor()
			for i, op := range q.para208.operations {
				var output interface{}
				switch op {
				case "Trie":
					output = nil
				case "insert":
					trie.Insert(q.para208.arguments[i][0])
					output = nil
				case "search":
					output = trie.Search(q.para208.arguments[i][0])
				case "startsWith":
					output = trie.StartsWith(q.para208.arguments[i][0])
				}
				expected := q.ans208.expected[i]
				if expected != nil && output != expected {
					t.Errorf("Step %d (%s): expected %v, got %v", i, op, expected, output)
				}
			}
			fmt.Printf("[input]: operations=%v       [output]: passed\n", q.para208.operations)
		})
	}
	fmt.Printf("\n\n\n")
}
