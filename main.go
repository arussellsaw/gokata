package main

import (
	"fmt"
)

func main() {
	var store transactionStore
	var statement statement
	var acc = account{store, statement}
	fmt.Println("Hello some unicode string i can't be bothered to write")
	acc.Deposit(transaction{"17/12/1993", 30})
	acc.Withdraw(transaction{"20/04/2004", 4})
	acc.Deposit(transaction{"14/04/2015", 2400})
	acc.Statement.write(acc.Log)
}
