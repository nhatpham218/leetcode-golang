package leetcode

import (
	"fmt"
	"testing"
)

type question0380 struct {
	name string
	para0380
	ans0380
}

// para is parameter
// one represents the first parameter
type para0380 struct {
	operations []string
	values     [][]int
}

// ans is answer
// one represents the first answer
type ans0380 struct {
	one []interface{}
}

func Test_Problem0380(t *testing.T) {

	qs := []question0380{
		{
			name: "example 1",
			para0380: para0380{
				operations: []string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"},
				values:     [][]int{{}, {1}, {2}, {2}, {}, {1}, {2}, {}},
			},
			ans0380: ans0380{[]interface{}{nil, true, false, true, 2, true, false, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0380------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			obj := Constructor()
			results := make([]interface{}, 0)
			for i, op := range q.operations {
				var result interface{}
				switch op {
				case "RandomizedSet":
					result = nil
				case "insert":
					result = obj.Insert(q.values[i][0])
				case "remove":
					result = obj.Remove(q.values[i][0])
				case "getRandom":
					result = obj.GetRandom()
				}
				results = append(results, result)
			}
			fmt.Printf("[input]: operations=%v, values=%v       [output]:%v\n", q.para0380.operations, q.para0380.values, results)
			// Note: getRandom returns random values, so we can't directly compare
			// We'll check insert and remove operations
			for i := 0; i < len(results); i++ {
				if q.operations[i] != "getRandom" && results[i] != q.ans0380.one[i] {
					t.Errorf("Operation %s at index %d: Expected %v, got %v", q.operations[i], i, q.ans0380.one[i], results[i])
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
