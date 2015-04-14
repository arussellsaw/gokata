package main

import ()

type transaction struct {
	timestamp string
	amount    int
}

type transactionStore struct {
	transactions []transaction
}
