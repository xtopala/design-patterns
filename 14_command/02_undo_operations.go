// Command - Undo Operations

// We can continue with the last example, and
// we can implement undo functionality.
// After we call the command we can foll back those
// changes that the command introduced.

// In order to do this, we need to modify a couple of things.
// So first, we'll add the Undo method to the command interface.

// But in addtion, one problem that we'll face is that
// some of the commands are actually going to fail.

package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Depositetd: ", amount, "\b, balance ist now ", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew: ", amount, "\b, balance ist now ", b.balance)
		return true
	}
	return false
}

// -> Like for an example here, when we withdraw a certain amount
// 	 we want to withdraw it only if the decreased amount is greater than
//	 or equal to the overdraft limit, otherwise the command does not apply.

// But, if the command doesn't apply that means you shouldn't be abble
// to undo it either, because that will leave the system in an unpredicable state.

// So we need to be able to have some sort of indicator
// whether the command succeeded or not.

// Another, question is: Where do we use this information?
// The answer is that we can use this whenever we actually
// perform the action.

// So when we call the command we can also store some information
// about whether or not the command actually succeeded.

type Command interface {
	Call()
	Undo()
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

// We'll add our new flag here.

type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
	}
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

// Now we can implement the Undo method.

func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

// ↑↑↑ We can not consider these operations symmetrically,
//	   but in our simple example it's going to work just fine.

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba)

	cmd2 := NewBankAccountCommand(&ba, Withdraw, 25)
	cmd2.Call()
	fmt.Println(ba)

	cmd2.Undo()
	fmt.Println(ba)
}
