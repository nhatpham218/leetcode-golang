package leetcode

import (
	"fmt"
	"testing"
)

type question3508 struct {
	name string
	para3508
	ans3508
}

// para is parameter
type para3508 struct {
	memoryLimit int
	ops         []string
	params      []interface{}
}

// ans is answer
type ans3508 struct {
	expected []interface{}
}

func Test_Problem3508(t *testing.T) {

	qs := []question3508{
		{
			name: "example 1",
			para3508: para3508{
				memoryLimit: 3,
				ops:         []string{"addPacket", "addPacket", "addPacket", "addPacket", "addPacket", "forwardPacket", "addPacket", "getCount"},
				params: []interface{}{
					[]interface{}{1, 4, 90},
					[]interface{}{2, 5, 90},
					[]interface{}{1, 4, 90},
					[]interface{}{3, 5, 95},
					[]interface{}{4, 5, 105},
					[]interface{}{},
					[]interface{}{5, 2, 110},
					[]interface{}{5, 100, 110},
				},
			},
			ans3508: ans3508{[]interface{}{true, true, false, true, true, []int{2, 5, 90}, true, 1}},
		},
		{
			name: "example 2",
			para3508: para3508{
				memoryLimit: 2,
				ops:         []string{"addPacket", "forwardPacket", "forwardPacket"},
				params: []interface{}{
					[]interface{}{7, 4, 90},
					[]interface{}{},
					[]interface{}{},
				},
			},
			ans3508: ans3508{[]interface{}{true, []int{7, 4, 90}, []int{}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3508------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3508, q.para3508
			obj := Constructor(p.memoryLimit)

			fmt.Printf("[input]: memoryLimit=%v\n", p.memoryLimit)

			for i, op := range p.ops {
				switch op {
				case "addPacket":
					params := p.params[i].([]interface{})
					result := obj.AddPacket(params[0].(int), params[1].(int), params[2].(int))
					fmt.Printf("[operation]: %s(%d, %d, %d) = %t\n", op, params[0].(int), params[1].(int), params[2].(int), result)
				case "forwardPacket":
					result := obj.ForwardPacket()
					fmt.Printf("[operation]: %s() = %v\n", op, result)
				case "getCount":
					params := p.params[i].([]interface{})
					result := obj.GetCount(params[0].(int), params[1].(int), params[2].(int))
					fmt.Printf("[operation]: %s(%d, %d, %d) = %d\n", op, params[0].(int), params[1].(int), params[2].(int), result)
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
