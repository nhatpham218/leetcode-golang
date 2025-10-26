package leetcode

// 2043. Simple Bank System
// https://leetcode.com/problems/simple-bank-system/
type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance}
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 < 1 || account1 > len(b.balance) || account2 < 1 || account2 > len(b.balance) {
		return false
	}
	if b.balance[account1-1] < money {
		return false
	}
	b.balance[account1-1] -= money
	b.balance[account2-1] += money
	return true
}

func (b *Bank) Deposit(account int, money int64) bool {
	if account < 1 || account > len(b.balance) {
		return false
	}
	b.balance[account-1] += money
	return true
}

func (b *Bank) Withdraw(account int, money int64) bool {
	if account < 1 || account > len(b.balance) {
		return false
	}
	if b.balance[account-1] < money {
		return false
	}
	b.balance[account-1] -= money
	return true
}
