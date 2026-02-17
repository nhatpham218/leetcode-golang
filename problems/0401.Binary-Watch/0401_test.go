package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

type question0401 struct {
	name string
	para0401
	ans0401
}

type para0401 struct {
	turnedOn int
}

type ans0401 struct {
	one []string
}

func Test_Problem0401(t *testing.T) {
	qs := []question0401{
		{
			name:     "example 1",
			para0401: para0401{turnedOn: 1},
			ans0401: ans0401{
				one: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
			},
		},
		{
			name:     "example 2",
			para0401: para0401{turnedOn: 9},
			ans0401:  ans0401{one: []string{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0401------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0401, q.para0401
			output := readBinaryWatch(p.turnedOn)
			fmt.Printf("[input]: turnedOn=%v       [output]:%v\n", p.turnedOn, output)
			// Answer may be in any order
			sort.Strings(output)
			expected := make([]string, len(q.ans0401.one))
			copy(expected, q.ans0401.one)
			sort.Strings(expected)
			if len(output) != len(expected) {
				t.Errorf("Expected %v, got %v", q.ans0401.one, output)
				return
			}
			for i := range output {
				if output[i] != expected[i] {
					t.Errorf("Expected %v, got %v", q.ans0401.one, output)
					return
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
