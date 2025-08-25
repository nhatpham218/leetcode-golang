package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3645 struct {
	name string
	para3645
	ans3645
}

// para is parameter
type para3645 struct {
	value []int
	limit []int
}

// ans is answer
type ans3645 struct {
	result int64
}

func Test_Problem3645(t *testing.T) {

	qs := []question3645{
		{
			name: "example 1",
			para3645: para3645{
				value: []int{3, 5, 8},
				limit: []int{2, 1, 3},
			},
			ans3645: ans3645{
				result: 16,
			},
		},
		{
			name: "example 2",
			para3645: para3645{
				value: []int{4, 2, 6},
				limit: []int{1, 1, 1},
			},
			ans3645: ans3645{
				result: 6,
			},
		},
		{
			name: "example 3",
			para3645: para3645{
				value: []int{4, 1, 5, 2},
				limit: []int{3, 3, 2, 3},
			},
			ans3645: ans3645{
				result: 12,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3645------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3645, q.para3645
			output := maxTotal(p.value, p.limit)
			fmt.Printf("[input]: value=%v, limit=%v       [output]:%v\n", p.value, p.limit, output)
			assert.Equal(t, q.ans3645.result, output, "maxTotal(%v, %v) should return %v", p.value, p.limit, q.ans3645.result)
		})
	}
	fmt.Printf("\n\n\n")
}
