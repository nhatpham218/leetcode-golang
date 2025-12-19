package leetcode

import (
	"fmt"
	"testing"
)

type question2092 struct {
	name string
	para2092
	ans2092
}

// para is parameter
// one represents the first parameter
type para2092 struct {
	n           int
	meetings    [][]int
	firstPerson int
}

// ans is answer
// one represents the first answer
type ans2092 struct {
	one []int
}

func Test_Problem2092(t *testing.T) {

	qs := []question2092{
		{
			name: "example 1",
			para2092: para2092{
				n:           6,
				meetings:    [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}},
				firstPerson: 1,
			},
			ans2092: ans2092{[]int{0, 1, 2, 3, 5}},
		},
		{
			name: "example 2",
			para2092: para2092{
				n:           4,
				meetings:    [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}},
				firstPerson: 3,
			},
			ans2092: ans2092{[]int{0, 1, 3}},
		},
		{
			name: "example 3",
			para2092: para2092{
				n:           5,
				meetings:    [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}},
				firstPerson: 1,
			},
			ans2092: ans2092{[]int{0, 1, 2, 3, 4}},
		},
		{
			name: "example 4",
			para2092: para2092{
				n:           5,
				meetings:    [][]int{{1, 4, 3}, {0, 4, 3}},
				firstPerson: 3,
			},
			ans2092: ans2092{[]int{0, 1, 3, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2092------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2092, q.para2092
			output := findAllPeople(p.n, p.meetings, p.firstPerson)
			fmt.Printf("[input]: n=%v, meetings=%v, firstPerson=%v       [output]:%v\n", p.n, p.meetings, p.firstPerson, output)
			if !equalSlicesUnordered(output, q.ans2092.one) {
				t.Errorf("Expected %v, got %v", q.ans2092.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

// equalSlicesUnordered checks if two slices contain the same elements (order doesn't matter)
func equalSlicesUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	// Count occurrences in both slices
	countA := make(map[int]int)
	countB := make(map[int]int)

	for _, v := range a {
		countA[v]++
	}
	for _, v := range b {
		countB[v]++
	}

	// Compare counts
	if len(countA) != len(countB) {
		return false
	}
	for k, v := range countA {
		if countB[k] != v {
			return false
		}
	}
	return true
}
