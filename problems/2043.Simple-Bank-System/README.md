# 2043. Simple Bank System

## Problem Description

You have been tasked with writing a program for a popular bank that will automate all its incoming transactions (transfer, deposit, and withdraw). The bank has n accounts numbered from 1 to n. The initial balance of each account is stored in a 0-indexed integer array balance, with the (i + 1)th account having an initial balance of balance[i].

Execute all the valid transactions. A transaction is valid if:

- The given account number(s) are between 1 and n, and
- The amount of money withdrawn or transferred from is less than or equal to the balance of the account.

Implement the Bank class:

- `Bank(long[] balance)` Initializes the object with the 0-indexed integer array balance.
- `boolean transfer(int account1, int account2, long money)` Transfers money dollars from the account numbered account1 to the account numbered account2. Return true if the transaction was successful, false otherwise.
- `boolean deposit(int account, long money)` Deposit money dollars into the account numbered account. Return true if the transaction was successful, false otherwise.
- `boolean withdraw(int account, long money)` Withdraw money dollars from the account numbered account. Return true if the transaction was successful, false otherwise.

## Examples

### Example 1:
```
Input
["Bank", "withdraw", "transfer", "deposit", "transfer", "withdraw"]
[[[10, 100, 20, 50, 30]], [3, 10], [5, 1, 20], [5, 20], [3, 4, 15], [10, 50]]
Output
[null, true, true, true, false, false]

Explanation
Bank bank = new Bank([10, 100, 20, 50, 30]);
bank.withdraw(3, 10);    // return true, account 3 has a balance of $20, so it is valid to withdraw $10.
                         // Account 3 has $20 - $10 = $10.
bank.transfer(5, 1, 20); // return true, account 5 has a balance of $30, so it is valid to transfer $20.
                         // Account 5 has $30 - $20 = $10, and account 1 has $10 + $20 = $30.
bank.deposit(5, 20);     // return true, it is valid to deposit $20 to account 5.
                         // Account 5 has $10 + $20 = $30.
bank.transfer(3, 4, 15); // return false, the current balance of account 3 is $10,
                         // so it is invalid to transfer $15 from it.
bank.withdraw(10, 50);   // return false, it is invalid because account 10 does not exist.
```

## Constraints

- `n == balance.length`
- `1 <= n, account, account1, account2 <= 10^5`
- `0 <= balance[i], money <= 10^12`
- At most `10^4` calls will be made to each function transfer, deposit, withdraw.

## Solution

### Algorithm Overview

This is a straightforward simulation problem where we need to maintain account balances and perform transactions with validation.

### Approach

1. **Data Structure**: Store the balances array in the Bank struct.
2. **Validation Helper**: Create a helper method to check if an account number is valid (between 1 and n).
3. **Withdraw**: Check if the account is valid and has sufficient balance, then deduct the amount.
4. **Deposit**: Check if the account is valid, then add the amount.
5. **Transfer**: Check if both accounts are valid and the source has sufficient balance, then perform the transfer.

### Key Points

- Account numbers are 1-indexed, but the balance array is 0-indexed
- All transactions must validate account numbers first
- Withdraw and transfer operations must check for sufficient balance
- Return true for successful transactions, false otherwise

### Complexity Analysis

- **Time Complexity**: O(1) for each operation
  - All operations involve constant time array access and arithmetic

- **Space Complexity**: O(n)
  - Store the balance array of size n

### Implementation

```go
type Bank struct {
}

func Constructor(balance []int64) Bank {
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
}

func (this *Bank) Deposit(account int, money int64) bool {
}

func (this *Bank) Withdraw(account int, money int64) bool {
}
```

### Test Cases

- Basic operations with valid accounts
- Invalid account numbers (out of range)
- Insufficient balance scenarios
- Multiple sequential operations affecting the same account

