package leetcode

import (
	"fmt"
	"testing"
)

type question3408 struct {
	name string
	para3408
	ans3408
}

// para is parameter
type para3408 struct {
	tasks  [][]int
	ops    []string
	params []interface{}
}

// ans is answer
type ans3408 struct {
	expected []interface{}
}

func Test_Problem3408(t *testing.T) {

	qs := []question3408{
		{
			name: "example 1",
			para3408: para3408{
				tasks:  [][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}},
				ops:    []string{"add", "edit", "execTop", "rmv", "add", "execTop"},
				params: []interface{}{[]interface{}{4, 104, 5}, []interface{}{102, 8}, nil, []interface{}{101}, []interface{}{5, 105, 15}, nil},
			},
			ans3408: ans3408{[]interface{}{nil, nil, 3, nil, nil, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3408------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3408, q.para3408
			obj := Constructor(p.tasks)

			fmt.Printf("[input]: tasks=%v\n", p.tasks)

			for i, op := range p.ops {
				switch op {
				case "add":
					params := p.params[i].([]interface{})
					obj.Add(params[0].(int), params[1].(int), params[2].(int))
					fmt.Printf("[operation]: %s(%d, %d, %d)\n", op, params[0].(int), params[1].(int), params[2].(int))
				case "edit":
					params := p.params[i].([]interface{})
					obj.Edit(params[0].(int), params[1].(int))
					fmt.Printf("[operation]: %s(%d, %d)\n", op, params[0].(int), params[1].(int))
				case "rmv":
					params := p.params[i].([]interface{})
					obj.Rmv(params[0].(int))
					fmt.Printf("[operation]: %s(%d)\n", op, params[0].(int))
				case "execTop":
					result := obj.ExecTop()
					fmt.Printf("[operation]: %s() = %d\n", op, result)
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
