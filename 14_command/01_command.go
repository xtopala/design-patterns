// Command

// Here, we're going to talk about a simple and familiar scenario.
// That scenario is going to be a bank account and a bank account
// is going to have withdrawal and deposit operations.

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

func (b *BankAccount) Withdraw(amount int) {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew: ", amount, "\b, balance ist now ", b.balance)
	}
}

// Now, we can try to setup the command pattern.
// There are different ways of actually handling commands:
// -> 1. The bank account itself handles the command
// -> 2. The command handles itself

// We're going to consider this particular scenarion where
// a command is just an interface.

type Command interface {
	Call()
}

// Now in terms of the actions that we can take upon the account
// there's gonna be two of them.

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account *BankAccount
	action  Action
	amount  int
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
	case Withdraw:
		b.account.Withdraw(b.amount)
	}
}

// We can also make a factory function for initializing
// all the different members.

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

// And now we can start usign this whole thing.

// Recap:
// -> This was the approach to implementing the Command pattern
//	  where the command sort of calls itself
// -> Command has some sort of Call method and the we specify
//	  the actual commands which implement the command interface

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba)

	cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
	cmd2.Call()
	fmt.Println(ba)
}
