package leetcode

import (
	"fmt"
	"testing"
)

type question3484 struct {
	name string
	para3484
	ans3484
}

// para is parameter
type para3484 struct {
	rows   int
	ops    []string
	params []interface{}
}

// ans is answer
type ans3484 struct {
	expected []interface{}
}

func Test_Problem3484(t *testing.T) {

	qs := []question3484{
		{
			name: "example 1",
			para3484: para3484{
				rows:   3,
				ops:    []string{"getValue", "setCell", "getValue", "setCell", "getValue", "resetCell", "getValue"},
				params: []interface{}{"=5+7", []interface{}{"A1", 10}, "=A1+6", []interface{}{"B2", 15}, "=A1+B2", []interface{}{"A1"}, "=A1+B2"},
			},
			ans3484: ans3484{[]interface{}{12, nil, 16, nil, 25, nil, 15}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3484------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3484, q.para3484
			obj := Constructor(p.rows)

			fmt.Printf("[input]: rows=%v\n", p.rows)

			for i, op := range p.ops {
				switch op {
				case "setCell":
					params := p.params[i].([]interface{})
					obj.SetCell(params[0].(string), params[1].(int))
					fmt.Printf("[operation]: %s(%s, %d)\n", op, params[0].(string), params[1].(int))
				case "resetCell":
					params := p.params[i].([]interface{})
					obj.ResetCell(params[0].(string))
					fmt.Printf("[operation]: %s(%s)\n", op, params[0].(string))
				case "getValue":
					formula := p.params[i].(string)
					result := obj.GetValue(formula)
					fmt.Printf("[operation]: %s(%s) = %d\n", op, formula, result)
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
