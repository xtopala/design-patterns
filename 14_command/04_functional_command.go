// Command - Functional Command

// There is this one thing we have to mention and
// that is that we can, strictly speaking, take a more
// functional approach to the way we create commands as well
// as the way we create composite commands.

package main

import "fmt"

// So here we'll have our Bank account, that we've used already,
// with our Deposit and Withdraw methods.
// Here, these methods are somewhat simpler than what we
// had previously, but it doesn't really matter now.

type BankAccount struct {
	Balance int
}

func Deposit(ba *BankAccount, amount int) {
	fmt.Println("Depositing: ", amount)
	ba.Balance += amount
}

func Withdraw(ba *BankAccount, amount int) {
	if ba.Balance >= amount {
		fmt.Println("Withdrawing: ", amount)
		ba.Balance -= amount
	}
}

// So if we imagine a bank account with a starting balance,
// that we define the commands uppo to.
// But, instead of defining them as separate structs which define
// the entire set of operations that need to be performed on a bank account,
// we can simply stick those operations inside a list of functions.

// So what does this approach gives us?
// Recap:
// -> Under one hand we lose all the information regarding
// 	  what kind of operation we were actually doing.
// -> In the previous examples when we had the command we could
// 	  save the command to a file or send it over the wire.
// 	  Here we're wrapping it into a function and so we're losing all
// 	  the information about what's actually going on.
// 	  Meaning we can't just go and save this function somewhere or
// 	  go into this function and look at what exactly is going on there.
// -> Strictly speaking we can have more than one thing happening inside
// 	  a single command.
// -> On the other hand, this is a more functional approach and it does
// 	  have it uses when we don't care about the structure and we just
// 	  wanto to put several invocations into a list, for that list to be
// 	  invoked or for the entire list to be undone, then a functional approach
// 	  is more viable.

func main() {
	ba := &BankAccount{0}
	var commands []func()

	commands = append(commands, func() {
		Deposit(ba, 100)
	})
	commands = append(commands, func() {
		Withdraw(ba, 25)
	})

	// Now that we have this list of functions,
	// we can just go through this array
	for _, cmd := range commands {
		cmd()
	}
	fmt.Println(ba)
}
