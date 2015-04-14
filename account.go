package main

import (
	"fmt"
)

type accountInterface interface {
	Deposit(transaction transaction) bool
	Withdraw(transaction transaction) bool
}

type account struct {
	Log       transactionStore
	Statement statementInterface
}

func (a *account) Deposit(transaction transaction) bool {
	if transaction.amount > 0 {
		fmt.Println("Valid transaction")
	} else {
		return false
	}
	a.Log.transactions = append(a.Log.transactions, transaction)
	return true
}

func (a *account) Withdraw(transaction transaction) bool {
	if transaction.amount > 0 {
		fmt.Println("Valid transaction")
	} else {
		return false
	}
	transaction.amount = 0 - transaction.amount
	a.Log.transactions = append(a.Log.transactions, transaction)
	return true
}
