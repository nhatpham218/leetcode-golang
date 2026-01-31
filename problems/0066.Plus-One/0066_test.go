package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question66 struct {
	name string
	para66
	ans66
}

type para66 struct {
	digits []int
}

type ans66 struct {
	one []int
}

func Test_Problem66(t *testing.T) {

	qs := []question66{
		{
			name:   "example 1",
			para66: para66{digits: []int{1, 2, 3}},
			ans66:  ans66{one: []int{1, 2, 4}},
		},
		{
			name:   "example 2",
			para66: para66{digits: []int{4, 3, 2, 1}},
			ans66:  ans66{one: []int{4, 3, 2, 2}},
		},
		{
			name:   "example 3",
			para66: para66{digits: []int{9}},
			ans66:  ans66{one: []int{1, 0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 66------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans66, q.para66
			output := plusOne(p.digits)
			fmt.Printf("[input]: digits=%v       [output]:%v\n", p.digits, output)
			if !reflect.DeepEqual(output, q.ans66.one) {
				t.Errorf("Expected %v, got %v", q.ans66.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

