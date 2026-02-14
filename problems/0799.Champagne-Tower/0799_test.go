package leetcode

import (
	"fmt"
	"math"
	"testing"
)

type question0799 struct {
	name string
	para0799
	ans0799
}

// para is parameter
// one represents the first parameter
type para0799 struct {
	poured     int
	query_row  int
	query_glass int
}

// ans is answer
// one represents the first answer
type ans0799 struct {
	one float64
}

func Test_Problem0799(t *testing.T) {

	qs := []question0799{
		{
			name: "example 1",
			para0799: para0799{
				poured:      1,
				query_row:   1,
				query_glass: 1,
			},
			ans0799: ans0799{0.00000},
		},
		{
			name: "example 2",
			para0799: para0799{
				poured:      2,
				query_row:   1,
				query_glass: 1,
			},
			ans0799: ans0799{0.50000},
		},
		{
			name: "example 3",
			para0799: para0799{
				poured:      100000009,
				query_row:   33,
				query_glass: 17,
			},
			ans0799: ans0799{1.00000},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0799------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0799, q.para0799
			output := champagneTower(p.poured, p.query_row, p.query_glass)
			fmt.Printf("[input]: poured=%v, query_row=%v, query_glass=%v       [output]:%v\n", p.poured, p.query_row, p.query_glass, output)
			if math.Abs(output-q.ans0799.one) > 1e-5 {
				t.Errorf("Expected %v, got %v", q.ans0799.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
