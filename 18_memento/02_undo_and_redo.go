// Memento - Undo and Redo

// One natural consequence of returning a Memento
// for every single change in a struct is that we effectively
// have the whole history of the system as it's modified.

// And as a consequence of that what we can do is
// we can implement Undo and Redo functionality as part
// of the Memento design patern.

// As soon as we have this history of events we can
// implement undo and redo because we have every single
// state preserved in a memento.

// Obviously, this is not always practical.
// Sometimes our memento will simply store to much information.

package main

import "fmt"

// But, in our world of bank accounts,
// we can actually set this kind of thing up.

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Memento
	current int // <- position indicator within changes
}

// We also need a constructor for Bank Accounts.

func NewBankAccount(ballance int) *BankAccount {
	b := &BankAccount{balance: ballance}
	b.changes = append(b.changes, &Memento{ballance})
	return b
}

func (b BankAccount) String() string {
	return fmt.Sprint("Ballance = $", b.balance, ", current = ", b.current)
}

// <- Just so we can print few things

// Let's have a Deposit operation, and we're
// not going to have Withdrawal operation but we'll
// have an operation for depositing money to the bank account
// and we'll see that it's rather more complicated than what we had before.

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	m := Memento{b.balance}
	b.changes = append(b.changes, &m)
	b.current++

	fmt.Println("Deposited", amount, ", balance is now", b.balance)
	return &m
}

// Ok, now since we've got the memento,
// we can obviously restore to that memento.

func (b *BankAccount) Restore(m *Memento) {
	if m != nil {
		b.balance = m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

// Now we can implement those fancy undo and redo operations.

func (b *BankAccount) Undo() *Memento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}

	return nil
}

func (b *BankAccount) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}

	return nil
}

// Let's put all of this together.

func main() {
	ba := NewBankAccount(100)
	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1", ba)
	ba.Undo()
	fmt.Println("Undo 2", ba)
	ba.Redo()
	fmt.Println("Redo", ba)
}
