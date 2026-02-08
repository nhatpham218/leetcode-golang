package leetcode

import (
	"fmt"
	"testing"
)

type question211 struct {
	name string
	para211
	ans211
}

type para211 struct {
	operations []string
	arguments  [][]string
}

type ans211 struct {
	expected []interface{}
}

func Test_Problem211(t *testing.T) {
	qs := []question211{
		{
			name: "example 1",
			para211: para211{
				operations: []string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
				arguments:  [][]string{{}, {"bad"}, {"dad"}, {"mad"}, {"pad"}, {"bad"}, {".ad"}, {"b.."}},
			},
			ans211: ans211{
				expected: []interface{}{nil, nil, nil, nil, false, true, true, true},
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 211------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			wd := Constructor211()
			for i, op := range q.para211.operations {
				var output interface{}
				switch op {
				case "WordDictionary":
					output = nil
				case "addWord":
					wd.AddWord(q.para211.arguments[i][0])
					output = nil
				case "search":
					output = wd.Search(q.para211.arguments[i][0])
				}
				expected := q.ans211.expected[i]
				if expected != nil && output != expected {
					t.Errorf("Step %d (%s): expected %v, got %v", i, op, expected, output)
				}
			}
			fmt.Printf("[input]: operations=%v       [output]: passed\n", q.para211.operations)
		})
	}
	fmt.Printf("\n\n\n")
}
