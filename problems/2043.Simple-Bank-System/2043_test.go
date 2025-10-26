package leetcode

import (
	"fmt"
	"testing"
)

type question2043 struct {
	name string
	para2043
	ans2043
}

type para2043 struct {
	balance  []int64
	commands []string
	args     [][]interface{}
}

type ans2043 struct {
	output []interface{}
}

func Test_Problem2043(t *testing.T) {

	qs := []question2043{}

	fmt.Printf("------------------------Leetcode Problem 2043------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans2043, q.para2043
			bank := Constructor(p.balance)

			var output []interface{}
			output = append(output, nil)

			for i := 1; i < len(p.commands); i++ {
				switch p.commands[i] {
				case "transfer":
					args := p.args[i]
					result := bank.Transfer(int(args[0].(int)), int(args[1].(int)), int64(args[2].(int)))
					output = append(output, result)
				case "deposit":
					args := p.args[i]
					result := bank.Deposit(int(args[0].(int)), int64(args[1].(int)))
					output = append(output, result)
				case "withdraw":
					args := p.args[i]
					result := bank.Withdraw(int(args[0].(int)), int64(args[1].(int)))
					output = append(output, result)
				}
			}

			fmt.Printf("[input]: %v       [output]:%v\n", p.commands, output)
			for i, o := range output {
				if i == 0 {
					continue
				}
				if o != a.output[i] {
					t.Errorf("expected %v, got %v", a.output[i], o)
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
