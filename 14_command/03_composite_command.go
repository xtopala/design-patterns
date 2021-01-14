// Command - Composite Command

// Now, let's say we want to transfer monet from one
// bank account to another bank account.
// We can consider this as a combination of two different commands.
// And one of those commands would withdraw the money from the first
// account, and the second one would deposit it into the second account.

// However, there are few limitations here.
// For example, they both should succeed or fail, meaning
// that if we fail to withdraw money from account, then we shouldn't
// be able to deposit this money somewhere else.
// That simply doesn't make sense.

// The whole transaction should fail.
// So we're going to implemet a -> Composite Command.
// This is a command which itself can consist of several other
// commands and it represents both the Command design pattern
// as well as the Composite Command pattern.

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

// Now the problem is that every single command that we
// operate on has to be able to have this success flag, it has
// to tell us whether it succeeded or not.

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(val bool)
}

// <- So we add this at the interface level.
// 	  We add together a getter and setter.

// Whenever there's a concrete command that implements this
// interface it actually has to somehow operate upon this idea of success.

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

// <- Down here, we once again need to implement the command interface.

func (b *BankAccountCommand) Succeeded() bool {
	return b.succeeded
}

func (b *BankAccountCommand) SetSucceeded(value bool) {
	b.succeeded = value
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

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

// Now we can go ahead and we can build a composite bank accounts command.
// And like we've said, a composite command is going to be just a
// collection of commands.

type CompositeBankAccountCommand struct {
	commands []Command
}

// <- We also need to implement the Command interface.

func (c CompositeBankAccountCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
	}
}

// We need to undo all commands in reverse order here.

func (c CompositeBankAccountCommand) Undo() {
	for i := range c.commands {
		c.commands[len(c.commands)-i-1].Undo()
	}
}

func (c CompositeBankAccountCommand) Succeeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}

func (c CompositeBankAccountCommand) SetSucceeded(value bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceeded(value)
	}
}

// So now that we have this Composite command, we can certainly start using it.
// And since we want to perform a transfer of money from one bank account
// to another, that action will aggregate the composite bank account command.

// We're going to compose it like this.

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

// Now here's the interesting thing, we need to make a factory function
// for this struct because it needs to be initialized correctly.

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{from: from, to: to, amount: amount}
	c.commands = append(c.commands, NewBankAccountCommand(from, Withdraw, amount))
	c.commands = append(c.commands, NewBankAccountCommand(to, Deposit, amount))

	return c
}

// Unfortunately, there is one more thing here we need to do.
// What happens here if the 1st operation fails?
// There's no sense to perform the next command, and certainly
// makes no sense at all to undo it later on, because as we said earlier
// either both commands succeed or they both fail.

// We need to redefine the Call method, which is going to be
// an alternative definition of the command invocation.

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Call()
			ok = cmd.Succeeded()
		} else {
			cmd.SetSucceeded(false)
		}
	}
}

// We can kind of do symmetrical operations, so our Undo will also work.

// Recap:
// -> If we have operations that we need to handle as a single operation
//	  we aggregate the composite command to get basically just the list
//	  of commands
// -> We also need to have a way of calling our operations
// -> Composite command pattern easily extends command pattern

func main() {
	from := BankAccount{100}
	to := BankAccount{0}

	mtc := NewMoneyTransferCommand(&from, &to, 25)
	mtc.Call()
	fmt.Println(from, to)

	mtc.Undo()
	fmt.Println(from, to)
}
