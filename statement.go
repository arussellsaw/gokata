package main

import (
	"fmt"
)

type statementInterface interface {
	write(transactionStore)
}

type statement struct{}

func (s statement) write(store transactionStore) {
	var balance int
	fmt.Println("DATE | AMOUNT | BALANCE")
	for i := 0; i < len(store.transactions); i++ {
		balance += store.transactions[i].amount
		fmt.Printf(
			"%s | %d | %d \n",
			store.transactions[i].timestamp,
			store.transactions[i].amount,
			balance)
	}
}
