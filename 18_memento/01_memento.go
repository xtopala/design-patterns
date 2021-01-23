// Memento - Memento

// So what we're going to take a look at now is
// a simpler implementation of the Memento design pattern.
// The idea of a Memento is that basically whenever we have a
// change in the system we can take a snapshot of the system and
// we can return that snapshot to the user, so they can sunsequently
// return the system to that state.

package main

import "fmt"

// Let's imagine we have a bank account.

type BankAccount struct {
	balance int
}

// And we want to allow the depositing of money.

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{b.balance}
}

// Well, hold on hold on.
// What is this -> Memento?
// Our Memento here is going to take a system snapshot
// and our system is rather simple, we only have the ballance
// of the system.

type Memento struct {
	Balance int
}

// And then of course, we need a mechanism for actually restoring
// the system to a state represented by the Memento.

func (b *BankAccount) Restore(m *Memento) {
	b.balance = m.Balance
}

// Let's take a look how all of this work.
// And this is great for restoring a system to previous state.

// There's only one problem though.
// Whenever we actually have this king of setup, we don't have
// a memento for the inital state of the bank account and if we
// want to have it then we'd have to improvise somehow.

// Because, remember if we make a typical method for a
// typical function for just initializing the bank account
// it's going to look something like this.

// *scrathced*
// func NewBankAccount(balance int) *BankAccount {
// return &BankAccount{balance: balance}
// }

// <- And the problem with this setup is that we basically return
//	  the bank account that we've constructed, but we're not returning
// 	  the Memento here.

// Of course, we could modify that, to something like this.

func NewBankAccount(balance int) (*BankAccount, *Memento) {
	return &BankAccount{balance: balance}, &Memento{balance}
}

// It's possible, to squeeze both of these.
// Nor particularly convinient though, but it's possible
// to return the initial memento

func main() {
	// ba := BankAccount{100}
	ba, m0 := NewBankAccount(100)
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)
	fmt.Println(ba)

	ba.Restore(m1)
	fmt.Println(ba)

	ba.Restore(m2)
	fmt.Println(ba)

	ba.Restore(m0)
	fmt.Println(ba)
}
