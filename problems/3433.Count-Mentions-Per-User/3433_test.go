package leetcode

import (
	"fmt"
	"testing"
)

type question3433 struct {
	name string
	para3433
	ans3433
}

// para is parameter
// one represents the first parameter
type para3433 struct {
	numberOfUsers int
	events        [][]string
}

// ans is answer
// one represents the first answer
type ans3433 struct {
	one []int
}

func Test_Problem3433(t *testing.T) {

	qs := []question3433{
		{
			name: "example 1",
			para3433: para3433{
				numberOfUsers: 2,
				events: [][]string{
					{"MESSAGE", "10", "id1 id0"},
					{"OFFLINE", "11", "0"},
					{"MESSAGE", "71", "HERE"},
				},
			},
			ans3433: ans3433{[]int{2, 2}},
		},
		{
			name: "example 2",
			para3433: para3433{
				numberOfUsers: 2,
				events: [][]string{
					{"MESSAGE", "10", "id1 id0"},
					{"OFFLINE", "11", "0"},
					{"MESSAGE", "12", "ALL"},
				},
			},
			ans3433: ans3433{[]int{2, 2}},
		},
		{
			name: "example 3",
			para3433: para3433{
				numberOfUsers: 2,
				events: [][]string{
					{"OFFLINE", "10", "0"},
					{"MESSAGE", "12", "HERE"},
				},
			},
			ans3433: ans3433{[]int{0, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3433------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3433, q.para3433
			output := countMentions(p.numberOfUsers, p.events)
			fmt.Printf("[input]: numberOfUsers=%v, events=%v       [output]:%v\n", p.numberOfUsers, p.events, output)
			if len(output) != len(q.ans3433.one) {
				t.Errorf("Expected length %v, got length %v", len(q.ans3433.one), len(output))
			}
			for i := range output {
				if output[i] != q.ans3433.one[i] {
					t.Errorf("Expected %v, got %v", q.ans3433.one, output)
					return
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
